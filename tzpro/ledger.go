// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type Ledger struct {
	Id             uint64    `json:"id"`
	Ledger         string    `json:"ledger"`
	TokenId        tezos.Z   `json:"token_id"`
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
	Supply         tezos.Z   `json:"total_supply"`
	VolMint        tezos.Z   `json:"total_minted"`
	VolBurn        tezos.Z   `json:"total_burned"`
	LastChange     int64     `json:"last_supply_change_block"`
	LastChangeTime time.Time `json:"last_supply_change_time"`
}

func (t Ledger) Address() tezos.Token {
	addr, _ := tezos.ParseAddress(t.Ledger)
	return tezos.NewToken(addr, t.TokenId)
}

type LedgerParams = Params[Ledger]

func NewLedgerParams() LedgerParams {
	return LedgerParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) GetLedger(ctx context.Context, addr tezos.Address, params LedgerParams) (*Ledger, error) {
	t := &Ledger{}
	u := params.WithPath(fmt.Sprintf("/v1/ledgers/%s", addr)).Url()
	if err := c.get(ctx, u, nil, t); err != nil {
		return nil, err
	}
	return t, nil
}

func (c *Client) ListLedgers(ctx context.Context, params LedgerParams) ([]*Ledger, error) {
	list := make([]*Ledger, 0)
	u := params.WithPath("/v1/ledgers").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
