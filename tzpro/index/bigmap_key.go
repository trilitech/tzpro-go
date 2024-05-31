// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/trilitech/tzpro-go/internal/util"
)

type MultiKey struct {
	named  map[string]any
	anon   []any
	single string
}

func DecodeMultiKey(key Key) (MultiKey, error) {
	mk := MultiKey{}
	buf, err := json.Marshal(key)
	if err != nil {
		return mk, err
	}
	err = json.Unmarshal(buf, &mk)
	return mk, err
}

func (k MultiKey) Len() int {
	if len(k.single) > 0 {
		return 1
	}
	return len(k.named) + len(k.anon)
}

func (k MultiKey) String() string {
	switch {
	case len(k.named) > 0:
		strs := make([]string, 0)
		for n, v := range k.named {
			strs = append(strs, fmt.Sprintf("%s=%s", n, v))
		}
		return strings.Join(strs, ",")
	case len(k.anon) > 0:
		strs := make([]string, 0)
		for _, v := range k.anon {
			strs = append(strs, util.ToString(v))
		}
		return strings.Join(strs, ",")
	default:
		return k.single
	}
}

func (k MultiKey) MarshalJSON() ([]byte, error) {
	switch {
	case len(k.named) > 0:
		return json.Marshal(k.named)
	case len(k.anon) > 0:
		return json.Marshal(k.anon)
	default:
		return []byte(strconv.Quote(k.single)), nil
	}
}

func (k *MultiKey) UnmarshalJSON(buf []byte) error {
	if len(buf) == 0 {
		return nil
	}
	switch buf[0] {
	case '{':
		m := make(map[string]any)
		if err := json.Unmarshal(buf, &m); err != nil {
			return err
		}
		k.named = m
	case '[':
		m := make([]any, 0)
		if err := json.Unmarshal(buf, &m); err != nil {
			return err
		}
		k.anon = m
	case '"':
		s, _ := strconv.Unquote(string(buf))
		k.single = s
	default:
		k.single = string(buf)
	}
	return nil
}

func (k MultiKey) GetString(path string) (string, bool) {
	return util.GetPathString(util.NonNil(k.named, k.anon, k.single), path)
}

func (k MultiKey) GetInt64(path string) (int64, bool) {
	return util.GetPathInt64(util.NonNil(k.named, k.anon, k.single), path)
}

func (k MultiKey) GetBig(path string) (*big.Int, bool) {
	return util.GetPathBig(util.NonNil(k.named, k.anon, k.single), path)
}

func (k MultiKey) GetTime(path string) (time.Time, bool) {
	return util.GetPathTime(util.NonNil(k.named, k.anon, k.single), path)
}

func (k MultiKey) GetAddress(path string) (Address, bool) {
	return util.GetPathAddress(util.NonNil(k.named, k.anon, k.single), path)
}

func (k MultiKey) GetValue(path string) (interface{}, bool) {
	return util.GetPathValue(util.NonNil(k.named, k.anon, k.single), path)
}

func (k MultiKey) Walk(path string, fn util.ValueWalkerFunc) error {
	val := util.NonNil(k.named, k.anon, k.single)
	if len(path) > 0 {
		var ok bool
		val, ok = util.GetPathValue(val, path)
		if !ok {
			return nil
		}
	}
	return util.WalkValueMap(path, val, fn)
}

func (k MultiKey) Unmarshal(val interface{}) error {
	buf, _ := json.Marshal(k)
	return json.Unmarshal(buf, val)
}
