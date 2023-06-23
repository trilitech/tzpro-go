// Copyright (c) 2013-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package util

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

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

var stringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

func ToString(t any) string {
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

func ShortDurationString(s string) string {
	return strings.TrimRight(strings.TrimRight(s, "0s"), "0m")
}

// func IsBase64(s string) bool {
//  _, err := base64.StdEncoding.DecodeString(s)
//  return err == nil
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
				b.WriteString(ToString(val.Index(i).Interface()))
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
