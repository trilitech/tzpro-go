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

	ListLedgerTokens(context.Context, Address, Params) ([]*Token, error)
	ListLedgerEvents(context.Context, Address, Params) ([]*TokenEvent, error)
	ListLedgerBalances(context.Context, Address, Params) ([]*TokenBalance, error)

	ListTokenEvents(context.Context, TokenAddress, Params) ([]*TokenEvent, error)
	ListTokenBalances(context.Context, TokenAddress, Params) ([]*TokenBalance, error)

	// firehose
	ListTokens(context.Context, Params) ([]*Token, error)
	ListLedgers(context.Context, Params) ([]*Ledger, error)
	ListMetadata(context.Context, Params) ([]*TokenMetadata, error)
}

func NewTokenAPI(c *client.Client) TokenAPI {
	return &tokenClient{client: c}
}

type tokenClient struct {
	client *client.Client
}

type Token struct {
	Id             uint64    `json:"id"`
	Ledger         string    `json:"ledger"`
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
}

func (t Token) Address() TokenAddress {
	addr, _ := ParseAddress(t.Ledger)
	return NewTokenAddress(addr, t.TokenId)
}

func (c *tokenClient) GetToken(ctx context.Context, addr TokenAddress) (*Token, error) {
	t := &Token{}
	u := fmt.Sprintf("/v1/tokens/%s", addr)
	if err := c.client.Get(ctx, u, nil, t); err != nil {
		return nil, err
	}
	return t, nil
}

func (c *tokenClient) ListTokens(ctx context.Context, params Params) ([]*Token, error) {
	list := make([]*Token, 0)
	u := params.WithPath("/v1/tokens").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
