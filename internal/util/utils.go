// Copyright (c) 2013-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package util

import (
	"encoding/binary"
	"encoding/hex"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func NonNil(vals ...interface{}) interface{} {
	for _, v := range vals {
		if v != nil {
			return v
		}
	}
	return nil
}

func ParseU64(s string) (u uint64) {
	buf, _ := hex.DecodeString(s)
	u = binary.BigEndian.Uint64(buf[:8])
	return
}
