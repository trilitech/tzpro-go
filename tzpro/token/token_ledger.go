// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package token

import (
	"context"
	"fmt"
	"time"
)

type Ledger struct {
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

func (t Ledger) Address() Address {
	addr, _ := ParseAddress(t.Ledger)
	return addr
}

func (c *tokenClient) GetLedger(ctx context.Context, addr Address) (*Ledger, error) {
	t := &Ledger{}
	u := fmt.Sprintf("/v1/ledgers/%s", addr)
	if err := c.client.Get(ctx, u, nil, t); err != nil {
		return nil, err
	}
	return t, nil
}

func (c *tokenClient) ListLedgers(ctx context.Context, params Query) ([]*Ledger, error) {
	list := make([]*Ledger, 0)
	u := params.WithPath("/v1/ledgers").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *tokenClient) ListLedgerTokens(ctx context.Context, addr Address, params Query) ([]*Token, error) {
	list := make([]*Token, 0)
	u := params.WithPath("/v1/ledgers/%s/tokens").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
