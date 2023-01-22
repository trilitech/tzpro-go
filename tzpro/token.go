// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type Token struct {
	Id             uint64        `json:"id"`
	Contract       tezos.Address `json:"contract"`
	TokenId        tezos.Z       `json:"token_id"`
	Kind           string        `json:"kind"`
	Type           string        `json:"type"`
	Name           string        `json:"name"`
	Symbol         string        `json:"symbol"`
	Decimals       int           `json:"decimals"`
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

type TokenParams struct {
	Params
}

func NewTokenParams() TokenParams {
	return TokenParams{NewParams()}
}

func (p TokenParams) WithLimit(v uint) TokenParams {
	p.Query.Set("limit", strconv.Itoa(int(v)))
	return p
}

func (p TokenParams) WithOffset(v uint) TokenParams {
	p.Query.Set("offset", strconv.Itoa(int(v)))
	return p
}

func (p TokenParams) WithCursor(v uint64) TokenParams {
	p.Query.Set("cursor", strconv.FormatUint(v, 10))
	return p
}

func (p TokenParams) WithOrder(o OrderType) TokenParams {
	p.Query.Set("order", string(o))
	return p
}

func (p TokenParams) WithDesc() TokenParams {
	p.Query.Set("order", string(OrderDesc))
	return p
}

func (p TokenParams) WithAsc() TokenParams {
	p.Query.Set("order", string(OrderAsc))
	return p
}

func (p TokenParams) WithContract(c tezos.Address) TokenParams {
	p.Query.Set("contract", c.String())
	return p
}

func (p TokenParams) WithTokenId(id tezos.Z) TokenParams {
	p.Query.Set("token_id", id.String())
	return p
}

func (p TokenParams) WithKind(k string) TokenParams {
	p.Query.Set("kind", k)
	return p
}

func (p TokenParams) WithType(t string) TokenParams {
	p.Query.Set("type", t)
	return p
}

func (p TokenParams) WithName(n string) TokenParams {
	p.Query.Set("name", n)
	return p
}

func (p TokenParams) WithSymbol(s string) TokenParams {
	p.Query.Set("symbol", s)
	return p
}

func (p TokenParams) WithDecimals(d int) TokenParams {
	p.Query.Set("decimals", strconv.Itoa(d))
	return p
}

func (p TokenParams) WithTags(t ...string) TokenParams {
	p.Query.Set("tags", strings.Join(t, ","))
	return p
}

func (c *Client) GetToken(ctx context.Context, addr tezos.Token, params TokenParams) (*Token, error) {
	t := &Token{}
	u := params.AppendQuery(fmt.Sprintf("/v1/tokens/%s", addr))
	if err := c.get(ctx, u, nil, t); err != nil {
		return nil, err
	}
	return t, nil
}

func (c *Client) ListTokens(ctx context.Context, params TokenParams) ([]*Token, error) {
	list := make([]*Token, 0)
	u := params.AppendQuery("/v1/tokens")
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
