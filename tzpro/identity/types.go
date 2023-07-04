// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package identity

import (
	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/internal/client"
	"blockwatch.cc/tzpro-go/internal/util"
)

type (
	Query = client.Query

	OpHash   = tezos.OpHash
	Address  = tezos.Address
	Token    = tezos.Token
	Z        = tezos.Z
	HexBytes = util.HexBytes
)

var (
	NewZ                = tezos.NewZ
	NewToken            = tezos.NewToken
	AddressTypeContract = tezos.AddressTypeContract
	ParseAddress        = tezos.ParseAddress
	MustParseAddress    = tezos.MustParseAddress
	NewAddress          = tezos.NewAddress
)
