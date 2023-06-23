// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/internal/client"
	"blockwatch.cc/tzpro-go/tzpro/token"
)

type (
	Params = client.Params

	OpHash  = tezos.OpHash
	Address = tezos.Address
	Z       = tezos.Z

	Token = token.Token
)

var (
	AddressTypeContract = tezos.AddressTypeContract
	ParseAddress        = tezos.ParseAddress
	NewAddress          = tezos.NewAddress
)
