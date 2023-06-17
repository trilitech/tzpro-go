// Copyright (c) 2013-2020 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type Time time.Time

func (f Time) Time() time.Time {
	return time.Time(f)
}

func (f *Time) UnmarshalText(buf []byte) error {
	val := string(buf)

	// try parsing as int
	if i, err := strconv.ParseInt(val, 10, 64); err == nil {
		// 1st try parsing as unix timestamp
		// detect UNIX timestamp scale: we choose somewhat arbitrarity
		// Dec 31, 9999 23:59:59 as cut-off time here
		switch {
		case i < 253402300799:
			// timestamp is in seconds
			*f = Time(time.Unix(i, 0).UTC())
		case i < 253402300799000:
			// timestamp is in milliseconds
			*f = Time(time.Unix(0, i*1000000).UTC())
		case i < 253402300799000000:
			// timestamp is in microseconds
			*f = Time(time.Unix(0, i*1000).UTC())
		default:
			// timestamp is in nanoseconds
			*f = Time(time.Unix(0, i).UTC())
		}
		return nil
	}
	t, err := time.Parse(time.RFC3339, val)
	*f = Time(t)
	return err
}

func (f *Time) UnmarshalJSON(data []byte) error {
	return f.UnmarshalText(bytes.Trim(data, "\""))
}

type Bool bool

func (b Bool) Bool() bool {
	return bool(b)
}

func (b Bool) True() bool {
	return bool(b)
}

func (b Bool) False() bool {
	return !bool(b)
}

func (b *Bool) UnmarshalText(buf []byte) error {
	switch s := string(buf); s {
	case "0", "false":
		*b = false
	case "1", "true":
		*b = true
	default:
		return fmt.Errorf("invalid bool value %q", s)
	}
	return nil
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	return b.UnmarshalText(bytes.Trim(data, "\""))
}

type StringList []string

func (l *StringList) UnmarshalText(data []byte) error {
	*l = strings.Split(string(data), ",")
	return nil
}

func (l *StringList) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		return l.UnmarshalText(bytes.Trim(data, `"`))
	}
	type alias *StringList
	return json.Unmarshal(data, alias(l))
}

var null = []byte(`null`)

var stringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

func toString(t any) string {
	switch t := t.(type) {
	case string:
		return t
	case time.Time:
		return t.Format(time.RFC3339)
	}
	val := reflect.Indirect(reflect.ValueOf(t))
	if !val.IsValid() {
		return ""
	}
	if val.Type().Implements(stringerType) {
		return t.(fmt.Stringer).String()
	}
	if s, err := toRawString(val.Interface()); err == nil {
		return s
	}
	return fmt.Sprintf("%v", val.Interface())
}

// func isBase64(s string) bool {
// 	_, err := base64.StdEncoding.DecodeString(s)
// 	return err == nil
// }

func toRawString(t any) (string, error) {
	val := reflect.Indirect(reflect.ValueOf(t))
	if !val.IsValid() {
		return "", nil
	}
	typ := val.Type()
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(val.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'g', -1, val.Type().Bits()), nil
	case reflect.String:
		return val.String(), nil
	case reflect.Bool:
		return strconv.FormatBool(val.Bool()), nil
	case reflect.Array:
		if typ.Elem().Kind() != reflect.Uint8 {
			break
		}
		// [...]byte
		var b []byte
		if val.CanAddr() {
			b = val.Slice(0, val.Len()).Bytes()
		} else {
			b = make([]byte, val.Len())
			reflect.Copy(reflect.ValueOf(b), val)
		}
		return hex.EncodeToString(b), nil
	case reflect.Slice:
		if typ.Elem().Kind() == reflect.Uint8 && !typ.Elem().Implements(stringerType) {
			// []byte
			b := val.Bytes()
			return hex.EncodeToString(b), nil
		} else {
			// anything else
			var b strings.Builder
			for i, l := 0, val.Len(); i < l; i++ {
				b.WriteString(toString(val.Index(i).Interface()))
				if i < l-1 {
					b.WriteByte(',')
				}
			}
			return b.String(), nil
		}
	case reflect.Map:
		return fmt.Sprintf("%#v", t), nil
	}
	return "", fmt.Errorf("no method for converting type %s (%v) to string", typ.String(), val.Kind())
}

type ValueWalkerFunc func(path string, value interface{}) error

func walkValueMap(name string, val interface{}, fn ValueWalkerFunc) error {
	switch t := val.(type) {
	case map[string]interface{}:
		if len(name) > 0 {
			name += "."
		}
		for n, v := range t {
			child := name + n
			if err := walkValueMap(child, v, fn); err != nil {
				return err
			}
		}
	case []interface{}:
		if len(name) > 0 {
			name += "."
		}
		for i, v := range t {
			child := name + strconv.Itoa(i)
			if err := walkValueMap(child, v, fn); err != nil {
				return err
			}
		}
	default:
		return fn(name, val)
	}
	return nil
}

func hasPath(val interface{}, path string) bool {
	if val == nil {
		return false
	}
	frag := strings.Split(path, ".")
	next := val
	for i, v := range frag {
		switch t := next.(type) {
		case map[string]interface{}:
			var ok bool
			next, ok = t[v]
			if !ok {
				return false
			}
		case []interface{}:
			idx, err := strconv.Atoi(v)
			if err != nil || len(t) < idx {
				return false
			}
			next = t[idx]
		default:
			return i == len(frag)-1
		}
	}
	return true
}

// Access nested map or array contents
func getPathString(val interface{}, path string) (string, bool) {
	if val == nil {
		return "", false
	}
	frag := strings.Split(path, ".")
	next := val
	for i, v := range frag {
		switch t := next.(type) {
		case map[string]interface{}:
			var ok bool
			next, ok = t[v]
			if !ok {
				return "", false
			}
		case []interface{}:
			idx, err := strconv.Atoi(v)
			if err != nil || len(t) < idx {
				return "", false
			}
			next = t[idx]
		default:
			return toString(next), i == len(frag)-1
		}
	}
	return toString(next), true
}

func getPathInt64(val interface{}, path string) (int64, bool) {
	str, ok := getPathString(val, path)
	if !ok {
		return 0, ok
	}
	i, err := strconv.ParseInt(str, 10, 64)
	return i, err == nil
}

func getPathBig(val interface{}, path string) (*big.Int, bool) {
	str, ok := getPathString(val, path)
	if !ok {
		return nil, ok
	}
	n := new(big.Int)
	_, err := fmt.Sscan(str, n)
	return n, err == nil
}

func getPathZ(val interface{}, path string) (tezos.Z, bool) {
	str, ok := getPathString(val, path)
	if !ok {
		return tezos.Zero, ok
	}
	n := new(big.Int)
	_, err := fmt.Sscan(str, n)
	return tezos.NewBigZ(n), err == nil
}

func getPathTime(val interface{}, path string) (time.Time, bool) {
	str, ok := getPathString(val, path)
	if !ok {
		return time.Time{}, ok
	}
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		// try parse as UNIX seconds
		if i, err2 := strconv.ParseInt(str, 10, 64); err2 == nil {
			t = time.Unix(i, 0)
			err = nil
		}
	}
	return t, err == nil
}

func getPathAddress(val interface{}, path string) (tezos.Address, bool) {
	str, ok := getPathString(val, path)
	if !ok {
		return tezos.InvalidAddress, ok
	}
	a, err := tezos.ParseAddress(str)
	return a, err == nil
}

func getPathValue(val interface{}, path string) (interface{}, bool) {
	if tree, ok := val.(map[string]interface{}); ok {
		frag := strings.Split(path, ".")
		for i, v := range frag {
			next, ok := tree[v]
			if !ok {
				return nil, false
			}
			switch t := next.(type) {
			case map[string]interface{}:
				tree = t
			default:
				return next, i == len(frag)-1
			}
		}
		return tree, true
	} else {
		return val, path == ""
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func nonNil(vals ...interface{}) interface{} {
	for _, v := range vals {
		if v != nil {
			return v
		}
	}
	return nil
}

func shortDurationString(s string) string {
	return strings.TrimRight(strings.TrimRight(s, "0s"), "0m")
}
