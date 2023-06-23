// Copyright (c) 2013-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package util

import (
	"bytes"
	"fmt"
)

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
