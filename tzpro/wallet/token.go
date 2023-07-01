// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package wallet

import (
	"context"
	"fmt"
)

func (c *walletClient) ListTokenBalances(ctx context.Context, addr Address, params Query) ([]*TokenBalance, error) {
	list := make([]*TokenBalance, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/balances", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *walletClient) ListTokenEvents(ctx context.Context, addr Address, params Query) ([]*TokenEvent, error) {
	list := make([]*TokenEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/token_events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
