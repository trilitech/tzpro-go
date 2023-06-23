// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"

	"blockwatch.cc/tzpro-go/internal/client"
)

type WalletAPI interface {
	ListTokenBalances(context.Context, Address, Params) ([]*TokenBalance, error)
	ListTokenEvents(context.Context, Address, Params) ([]*TokenEvent, error)

	ListDexEvents(context.Context, Address, Params) ([]*DexEvent, error)
	ListDexPositions(context.Context, Address, Params) ([]*DexPosition, error)
	ListDexTrades(context.Context, Address, Params) ([]*DexTrade, error)

	ListFarmEvents(context.Context, Address, Params) ([]*FarmEvent, error)
	ListFarmPositions(context.Context, Address, Params) ([]*FarmPosition, error)

	ListLendingEvents(context.Context, Address, Params) ([]*LendingEvent, error)
	ListLendingPositions(context.Context, Address, Params) ([]*LendingPosition, error)

	ListNftEvents(context.Context, Address, Params) ([]*NftEvent, error)
	ListNftPositions(context.Context, Address, Params) ([]*NftPosition, error)
	ListNftTrades(context.Context, Address, Params) ([]*NftTrade, error)
}

func NewWalletAPI(c *client.Client) WalletAPI {
	return &walletClient{client: c}
}

type walletClient struct {
	client *client.Client
}
