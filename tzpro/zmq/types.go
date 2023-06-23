// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package zmq

import (
	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/tzpro/index"
)

type (
	OpHash    = tezos.OpHash
	BlockHash = tezos.BlockHash
	Op        = index.Op
	Block     = index.Block
	Status    = index.Status
)

var (
	ParseOpHash    = tezos.ParseOpHash
	ParseBlockHash = tezos.ParseBlockHash
)
