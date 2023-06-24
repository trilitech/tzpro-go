// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package token

import (
	"context"
	"fmt"
)

type TokenBalance struct {
	Id         int64   `json:"id"`
	Owner      Address `json:"owner"`
	Ledger     Address `json:"ledger"`
	TokenId    Z       `json:"token_id"`
	Kind       string  `json:"token_kind"`
	Type       string  `json:"token_type"`
	Name       string  `json:"name"`
	Symbol     string  `json:"symbol"`
	Decimals   int     `json:"decimals"`
	Balance    Z       `json:"balance"`
	FirstBlock int64   `json:"first_block"`
	LastBlock  int64   `json:"last_block"`
	NTransfers int     `json:"num_transfers"`
	NMints     int     `json:"num_mints"`
	NBurns     int     `json:"num_burns"`
	VolSent    Z       `json:"vol_sent"`
	VolRecv    Z       `json:"vol_recv"`
	VolMint    Z       `json:"vol_mint"`
	VolBurn    Z       `json:"vol_burn"`
}

func (c *tokenClient) ListLedgerBalances(ctx context.Context, addr Address, params Query) ([]*TokenBalance, error) {
	list := make([]*TokenBalance, 0)
	u := params.WithPath(fmt.Sprintf("/v1/ledgers/%s/balances", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *tokenClient) ListTokenBalances(ctx context.Context, addr TokenAddress, params Query) ([]*TokenBalance, error) {
	list := make([]*TokenBalance, 0)
	u := params.WithPath(fmt.Sprintf("/v1/tokens/%s/balances", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
