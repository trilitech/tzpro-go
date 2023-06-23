// Copyright (c) 2013-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package util

import (
	"bytes"
	"strconv"
	"time"
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
