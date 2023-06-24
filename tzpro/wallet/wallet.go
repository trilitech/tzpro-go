// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"

	"blockwatch.cc/tzpro-go/internal/client"
)

type WalletAPI interface {
	ListTokenBalances(context.Context, Address, Query) ([]*TokenBalance, error)
	ListTokenEvents(context.Context, Address, Query) ([]*TokenEvent, error)

	ListDexEvents(context.Context, Address, Query) ([]*DexEvent, error)
	ListDexPositions(context.Context, Address, Query) ([]*DexPosition, error)
	ListDexTrades(context.Context, Address, Query) ([]*DexTrade, error)

	ListFarmEvents(context.Context, Address, Query) ([]*FarmEvent, error)
	ListFarmPositions(context.Context, Address, Query) ([]*FarmPosition, error)

	ListLendingEvents(context.Context, Address, Query) ([]*LendingEvent, error)
	ListLendingPositions(context.Context, Address, Query) ([]*LendingPosition, error)

	ListNftEvents(context.Context, Address, Query) ([]*NftEvent, error)
	ListNftPositions(context.Context, Address, Query) ([]*NftPosition, error)
	ListNftTrades(context.Context, Address, Query) ([]*NftTrade, error)
}

func NewWalletAPI(c *client.Client) WalletAPI {
	return &walletClient{client: c}
}

type walletClient struct {
	client *client.Client
}
