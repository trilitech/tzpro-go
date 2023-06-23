// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package client

import (
	"encoding"
	"encoding/hex"
	"fmt"
	"hash/fnv"

	"bytes"
	"encoding/json"
	"reflect"
	"sync"

	"blockwatch.cc/tzpro-go/internal/util"
)

type Decoder struct {
	id    uint32
	idx   []int // we only handle flat structs because thats what the SDK uses
	flags []int
}

func DecodeSlice(buf []byte, fields []string, val any) error {
	// val must be pointer to slice
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("decode: non-pointer passed to DecodeSlice")
	}
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("decode: non slice type %T for DecodeSlice", val)
	}

	etyp := v.Type().Elem()
	dec, err := buildDecoder(etyp, fields)
	if err != nil {
		return err
	}

	jdec := json.NewDecoder(bytes.NewReader(buf))
	_, err = jdec.Token()
	if err != nil {
		return err
	}

	// walk outer json array [
	for jdec.More() {
		elem := reflect.New(etyp)
		ev := elem
		if elem.Elem().Kind() == reflect.Ptr {
			ev.Elem().Set(reflect.New(elem.Elem().Type().Elem()))
			ev = reflect.Indirect(elem)
		}
		err = dec.decode(jdec, ev)
		if err != nil {
			return err
		}
		v.Set(reflect.Append(v, elem.Elem()))
	}

	// consume outer json arry closing bracket ]
	_, err = jdec.Token()
	return err
}

func Decode(buf []byte, fields []string, val any) error {
	// val must be pointer to struct
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("decode: non-pointer passed to Decode")
	}
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("decode: non slice type %T for Decode", val)
	}
	dec, err := buildDecoder(v.Type(), fields)
	if err != nil {
		return err
	}

	jdec := json.NewDecoder(bytes.NewReader(buf))
	return dec.decode(jdec, v)
}

func (d *Decoder) decode(dec *json.Decoder, dst reflect.Value) error {
	// read open bracket
	_, err := dec.Token()
	if err != nil {
		return err
	}

	// allocate and deref ptr types
	dst = derefValue(dst)

	// while the array contains values
	for i, pos := range d.idx {
		// custom pre-decoding
		f := derefValue(dst.Field(pos))
		switch {
		case d.flags[i]&fieldFlagHex > 0:
			// hex: decode hex to bin, then call binary unmarshaler
			var s string
			if err := dec.Decode(&s); err != nil {
				return err
			}
			buf, err := hex.DecodeString(s)
			if err != nil {
				return err
			}
			if err := f.Addr().Interface().(encoding.BinaryUnmarshaler).UnmarshalBinary(buf); err != nil {
				return err
			}
		case d.flags[i]&fieldFlagTime > 0:
			// time: decode int or time string
			var tm util.Time
			if err := dec.Decode(&tm); err != nil {
				return err
			}
			f.Set(reflect.ValueOf(tm.Time()))
		case d.flags[i]&fieldFlagBool > 0:
			// bool: decode int or string
			var b util.Bool
			if err := dec.Decode(&b); err != nil {
				return err
			}
			f.Set(reflect.ValueOf(b.Bool()))
		default:
			// decode an array value
			if err := dec.Decode(f.Addr().Interface()); err != nil {
				return err
			}
		}
	}

	// read closing bracket
	_, err = dec.Token()
	return err
}

var decoderMap = make(map[uint32]*Decoder)
var decoderLock sync.RWMutex

func typeHash(n string, s ...string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(n))
	for _, v := range s {
		h.Write([]byte(v))
	}
	return h.Sum32()
}

func buildDecoder(typ reflect.Type, fields []string) (*Decoder, error) {
	key := typeHash(typ.Name(), fields...)
	decoderLock.RLock()
	d, ok := decoderMap[key]
	decoderLock.RUnlock()
	if ok {
		return d, nil
	}
	tinfo, err := getReflectTypeInfo(typ, "json")
	if err != nil {
		return nil, err
	}
	if len(fields) == 0 {
		fields = tinfo.Aliases()
	}
	d = &Decoder{
		id:    key,
		idx:   make([]int, len(fields)),
		flags: make([]int, len(fields)),
	}

	for i, f := range fields {
		fi, ok := tinfo.Find(f)
		if !ok {
			return nil, fmt.Errorf("decode: missing type field %q", f)
		}
		// skip ignore fields
		if fi.ContainsFlag(fieldFlagIgnore) {
			continue
		}
		d.idx[i] = fi.Idx[0] // first index only, no nested structs
		d.flags[i] = fi.Flags
	}
	decoderLock.Lock()
	decoderMap[key] = d
	decoderLock.Unlock()
	return d, nil
}
