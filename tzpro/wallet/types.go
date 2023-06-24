// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package wallet

import (
	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/internal/client"
	"blockwatch.cc/tzpro-go/tzpro/defi"
	"blockwatch.cc/tzpro-go/tzpro/nft"
	"blockwatch.cc/tzpro-go/tzpro/token"
)

type (
	Query = client.Query

	OpHash  = tezos.OpHash
	Address = tezos.Address
	Token   = tezos.Token
	Z       = tezos.Z

	TokenEvent   = token.TokenEvent
	TokenBalance = token.TokenBalance

	DexEvent    = defi.DexEvent
	DexPosition = defi.DexPosition
	DexTrade    = defi.DexTrade

	FarmEvent    = defi.FarmEvent
	FarmPosition = defi.FarmPosition

	LendingEvent    = defi.LendingEvent
	LendingPosition = defi.LendingPosition

	NftEvent    = nft.NftEvent
	NftPosition = nft.NftPosition
	NftTrade    = nft.NftTrade
)

var (
	NewQuery = client.NewQuery
)
