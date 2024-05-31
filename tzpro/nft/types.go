// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package nft

import (
	"github.com/trilitech/tzgo/tezos"
	"github.com/trilitech/tzpro-go/internal/client"
	"github.com/trilitech/tzpro-go/tzpro/token"
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
