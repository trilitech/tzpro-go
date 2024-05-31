// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package token

import (
	"github.com/trilitech/tzgo/tezos"
	"github.com/trilitech/tzpro-go/internal/client"
)

type (
	Query = client.Query

	OpHash       = tezos.OpHash
	Address      = tezos.Address
	TokenAddress = tezos.Token
	Z            = tezos.Z
)

var (
	ParseAddress    = tezos.ParseAddress
	NewTokenAddress = tezos.NewToken
)
