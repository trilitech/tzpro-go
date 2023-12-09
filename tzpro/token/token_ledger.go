// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package token

import (
	"context"
	"fmt"
	"time"
)

type Ledger struct {
	Id          uint64    `json:"id"`
	Contract    Address   `json:"contract"`
	Creator     Address   `json:"creator"`
	Category    string    `json:"category"`
	Kind        string    `json:"token_kind"`
	Type        string    `json:"token_type"`
	FirstBlock  int64     `json:"first_block"`
	FirstTime   time.Time `json:"first_time"`
	NumUsers    int       `json:"num_users"`
	NumHolders  int       `json:"num_holders"`
	NumTokens   int       `json:"num_tokens"`
	Tags        []string  `json:"tags"`
	CodeHash    string    `json:"code_hash"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Homepage    string    `json:"homepage"`
	Version     string    `json:"version"`
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
	u := params.WithPath(fmt.Sprintf("/v1/ledgers/%s/tokens", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
