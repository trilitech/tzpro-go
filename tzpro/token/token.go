// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package token

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzpro-go/internal/client"
)

type TokenAPI interface {
	GetToken(context.Context, TokenAddress) (*Token, error)
	GetLedger(context.Context, Address) (*Ledger, error)

	GetTokenMetadata(context.Context, TokenAddress) (*TokenMetadata, error)
	GetLedgerMetadata(context.Context, Address) (*TokenMetadata, error)

	ListLedgerTokens(context.Context, Address, Query) ([]*Token, error)
	ListLedgerEvents(context.Context, Address, Query) ([]*TokenEvent, error)
	ListLedgerBalances(context.Context, Address, Query) ([]*TokenBalance, error)

	ListTokenEvents(context.Context, TokenAddress, Query) ([]*TokenEvent, error)
	ListTokenBalances(context.Context, TokenAddress, Query) ([]*TokenBalance, error)

	// firehose
	ListTokens(context.Context, Query) ([]*Token, error)
	ListEvents(context.Context, Query) ([]*TokenEvent, error)
	ListLedgers(context.Context, Query) ([]*Ledger, error)
	ListMetadata(context.Context, Query) ([]*TokenMetadata, error)
}

func NewTokenAPI(c *client.Client) TokenAPI {
	return &tokenClient{client: c}
}

type tokenClient struct {
	client *client.Client
}

type Token struct {
	Id             uint64    `json:"id"`
	Contract       Address   `json:"contract"`
	TokenId        Z         `json:"token_id"`
	Kind           string    `json:"token_kind"`
	Type           string    `json:"token_type"`
	Category       string    `json:"category"`
	Name           string    `json:"name"`
	Symbol         string    `json:"symbol"`
	Decimals       int       `json:"decimals"`
	Logo           string    `json:"logo"`
	Tags           []string  `json:"tags"`
	Creator        string    `json:"creator"`
	FirstBlock     int64     `json:"first_block"`
	FirstTime      time.Time `json:"first_time"`
	Supply         Z         `json:"total_supply"`
	VolMint        Z         `json:"total_minted"`
	VolBurn        Z         `json:"total_burned"`
	LastChange     int64     `json:"last_supply_change_block"`
	LastChangeTime time.Time `json:"last_supply_change_time"`
	PriceUSD       float64   `json:"price_usd,string"`
	McapUSD        float64   `json:"mcap_usd,string"`
}

func (t Token) Address() TokenAddress {
	return NewTokenAddress(t.Contract, t.TokenId)
}

func (c *tokenClient) GetToken(ctx context.Context, addr TokenAddress) (*Token, error) {
	t := &Token{}
	u := fmt.Sprintf("/v1/tokens/%s", addr)
	if err := c.client.Get(ctx, u, nil, t); err != nil {
		return nil, err
	}
	return t, nil
}

func (c *tokenClient) ListTokens(ctx context.Context, params Query) ([]*Token, error) {
	list := make([]*Token, 0)
	u := params.WithPath("/v1/tokens").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
