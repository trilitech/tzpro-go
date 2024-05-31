// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package wallet

import (
	"context"

	"github.com/trilitech/tzpro-go/internal/client"
)

type WalletAPI interface {
	// Token API
	ListTokenBalances(context.Context, Address, Query) ([]*TokenBalance, error)
	ListTokenEvents(context.Context, Address, Query) ([]*TokenEvent, error)

	// DEX API
	ListDexEvents(context.Context, Address, Query) ([]*DexEvent, error)
	ListDexPositions(context.Context, Address, Query) ([]*DexPosition, error)
	ListDexTrades(context.Context, Address, Query) ([]*DexTrade, error)

	// Farm API
	ListFarmEvents(context.Context, Address, Query) ([]*FarmEvent, error)
	ListFarmPositions(context.Context, Address, Query) ([]*FarmPosition, error)

	// Lending API
	ListLendingEvents(context.Context, Address, Query) ([]*LendingEvent, error)
	ListLendingPositions(context.Context, Address, Query) ([]*LendingPosition, error)

	// NFT API
	ListNftEvents(context.Context, Address, Query) ([]*NftEvent, error)
	ListNftPositions(context.Context, Address, Query) ([]*NftPosition, error)
	ListNftTrades(context.Context, Address, Query) ([]*NftTrade, error)

	// Identity API
	ListDomains(context.Context, Address, Query) ([]*Domain, error)
	ListDomainEvents(context.Context, Address, Query) ([]*DomainEvent, error)
	GetProfile(context.Context, Address) (*Profile, error)
	ListProfileEvents(context.Context, Address, Query) ([]*ProfileEvent, error)
	ListProfileClaims(context.Context, Address, Query) ([]*ProfileClaim, error)
}

func NewWalletAPI(c *client.Client) WalletAPI {
	return &walletClient{client: c}
}

type walletClient struct {
	client *client.Client
}
