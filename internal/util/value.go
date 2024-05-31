// Copyright (c) 2013-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package util

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/trilitech/tzgo/tezos"
)

type ValueWalkerFunc func(path string, value interface{}) error

func WalkValueMap(name string, val interface{}, fn ValueWalkerFunc) error {
	switch t := val.(type) {
	case map[string]interface{}:
		if len(name) > 0 {
			name += "."
		}
		for n, v := range t {
			child := name + n
			if err := WalkValueMap(child, v, fn); err != nil {
				return err
			}
		}
	case []interface{}:
		if len(name) > 0 {
			name += "."
		}
		for i, v := range t {
			child := name + strconv.Itoa(i)
			if err := WalkValueMap(child, v, fn); err != nil {
				return err
			}
		}
	default:
		return fn(name, val)
	}
	return nil
}

func HasPath(val interface{}, path string) bool {
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
func GetPathString(val interface{}, path string) (string, bool) {
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
			return ToString(next), i == len(frag)-1
		}
	}
	return ToString(next), true
}

func GetPathInt64(val interface{}, path string) (int64, bool) {
	str, ok := GetPathString(val, path)
	if !ok {
		return 0, ok
	}
	i, err := strconv.ParseInt(str, 10, 64)
	return i, err == nil
}

func GetPathBig(val interface{}, path string) (*big.Int, bool) {
	str, ok := GetPathString(val, path)
	if !ok {
		return nil, ok
	}
	n := new(big.Int)
	_, err := fmt.Sscan(str, n)
	return n, err == nil
}

func GetPathZ(val interface{}, path string) (tezos.Z, bool) {
	str, ok := GetPathString(val, path)
	if !ok {
		return tezos.Zero, ok
	}
	n := new(big.Int)
	_, err := fmt.Sscan(str, n)
	return tezos.NewBigZ(n), err == nil
}

func GetPathTime(val interface{}, path string) (time.Time, bool) {
	str, ok := GetPathString(val, path)
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

func GetPathAddress(val interface{}, path string) (tezos.Address, bool) {
	str, ok := GetPathString(val, path)
	if !ok {
		return tezos.InvalidAddress, ok
	}
	a, err := tezos.ParseAddress(str)
	return a, err == nil
}

func GetPathValue(val interface{}, path string) (interface{}, bool) {
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
