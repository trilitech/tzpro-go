// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package token

import (
	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/internal/client"
)

type (
	Params = client.Params

	OpHash       = tezos.OpHash
	Address      = tezos.Address
	TokenAddress = tezos.Token
	Z            = tezos.Z
)

var (
	ParseAddress    = tezos.ParseAddress
	NewTokenAddress = tezos.NewToken
)
