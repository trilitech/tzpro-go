// Copyright (c) 2013-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package util

import (
	"encoding/hex"
	"io"
)

// HexBytes represents bytes as a JSON string of hexadecimal digits
type HexBytes []byte

// UnmarshalText umarshals a hex string to bytes. It implements the
// encoding.TextUnmarshaler interface, so that HexBytes can be used in Go
// structs in combination with the standard JSON library.
func (h *HexBytes) UnmarshalText(data []byte) error {
	dst := make([]byte, hex.DecodedLen(len(data)))
	if _, err := hex.Decode(dst, data); err != nil {
		return err
	}
	*h = dst
	return nil
}

// ReadBytes copies size bytes from r into h. If h is nil or too short,
// a new byte slice is allocated. Fails with io.ErrShortBuffer when
// less than size bytes can be read from r.
func (h *HexBytes) ReadBytes(r io.Reader, size int) error {
	if cap(*h) < size {
		*h = make([]byte, size)
	} else {
		*h = (*h)[:size]
	}
	n, err := r.Read(*h)
	if err != nil {
		return err
	}
	if n < size {
		return io.ErrShortBuffer
	}
	return nil
}

// MarshalText converts a byte slice to a hex encoded string. It implements the
// encoding.TextMarshaler interface, so that HexBytes can be used in Go
// structs in combination with the standard JSON library.
func (h HexBytes) MarshalText() ([]byte, error) {
	return []byte(hex.EncodeToString(h)), nil
}

// String converts a byte slice to a hex encoded string
func (h HexBytes) String() string {
	return hex.EncodeToString(h)
}

// Bytes type-casts HexBytes back to a byte slice
func (h HexBytes) Bytes() []byte {
	return []byte(h)
}
