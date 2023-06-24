// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package nft

import (
	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/internal/client"
	"blockwatch.cc/tzpro-go/tzpro/token"
)

type (
	Query = client.Query

	OpHash  = tezos.OpHash
	Address = tezos.Address
	Z       = tezos.Z

	Token = token.Token
)

var (
	ParseAddress = tezos.ParseAddress
)
