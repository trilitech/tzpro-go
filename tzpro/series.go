// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"time"
)

type FillMode string

const (
	FillModeInvalid FillMode = ""
	FillModeNone    FillMode = "none"
	FillModeNull    FillMode = "null"
	FillModeLast    FillMode = "last"
	FillModeLinear  FillMode = "linear"
	FillModeZero    FillMode = "zero"
)

const (
	Collapse1m = time.Minute
	Collapse1h = time.Hour
	Collapse1d = 24 * time.Hour
	Collapse1w = 7 * 24 * time.Hour
)
