// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"

	"blockwatch.cc/tzgo/tezos"
)

type TokenBalance struct {
	Id         int64         `json:"id"`
	Owner      tezos.Address `json:"owner"`
	Contract   tezos.Address `json:"contract"`
	TokenId    tezos.Z       `json:"token_id"`
	Kind       string        `json:"kind"`
	Type       string        `json:"type"`
	Name       string        `json:"name"`
	Symbol     string        `json:"symbol"`
	Decimals   int           `json:"decimals"`
	Balance    tezos.Z       `json:"balance"`
	FirstBlock int64         `json:"first_block"`
	LastBlock  int64         `json:"last_block"`
	NTransfers int           `json:"num_transfers"`
	NMints     int           `json:"num_mints"`
	NBurns     int           `json:"num_burns"`
	VolSent    tezos.Z       `json:"vol_sent"`
	VolRecv    tezos.Z       `json:"vol_recv"`
	VolMint    tezos.Z       `json:"vol_mint"`
	VolBurn    tezos.Z       `json:"vol_burn"`
}

func (c *Client) ListTokenOwners(ctx context.Context, addr tezos.Token, params TokenParams) ([]*TokenBalance, error) {
	list := make([]*TokenBalance, 0)
	u := params.WithPath(fmt.Sprintf("/v1/tokens/%s/owners", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListWalletBalances(ctx context.Context, addr tezos.Address, params TokenParams) ([]*TokenBalance, error) {
	list := make([]*TokenBalance, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/token/balances", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
