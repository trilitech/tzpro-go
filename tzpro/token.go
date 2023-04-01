// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type Token struct {
	Id             uint64        `json:"id"`
	Contract       tezos.Address `json:"contract"`
	TokenId        tezos.Z       `json:"token_id"`
	Kind           string        `json:"kind"`
	Type           string        `json:"type"`
	Category       string        `json:"category"`
	Name           string        `json:"name"`
	Symbol         string        `json:"symbol"`
	Decimals       int           `json:"decimals"`
	Logo           string        `json:"logo"`
	Tags           []string      `json:"tags"`
	Creator        tezos.Address `json:"creator"`
	FirstBlock     int64         `json:"first_block"`
	FirstTime      time.Time     `json:"first_time"`
	Supply         tezos.Z       `json:"total_supply"`
	VolMint        tezos.Z       `json:"total_minted"`
	VolBurn        tezos.Z       `json:"total_burned"`
	LastChange     int64         `json:"last_supply_change_block"`
	LastChangeTime time.Time     `json:"last_supply_change_time"`
}

func (t Token) Address() tezos.Token {
	return tezos.NewToken(t.Contract, t.TokenId)
}

type TokenParams = Params[Token]

func NewTokenParams() TokenParams {
	return TokenParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) GetToken(ctx context.Context, addr tezos.Token, params TokenParams) (*Token, error) {
	t := &Token{}
	u := params.WithPath(fmt.Sprintf("/v1/tokens/%s", addr)).Url()
	if err := c.get(ctx, u, nil, t); err != nil {
		return nil, err
	}
	return t, nil
}

func (c *Client) ListTokens(ctx context.Context, params TokenParams) ([]*Token, error) {
	list := make([]*Token, 0)
	u := params.WithPath("/v1/tokens").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
