// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package wallet

import (
	"context"
	"fmt"
)

func (c *walletClient) ListDexEvents(ctx context.Context, addr Address, params Query) ([]*DexEvent, error) {
	list := make([]*DexEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/dex_events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *walletClient) ListDexPositions(ctx context.Context, addr Address, params Query) ([]*DexPosition, error) {
	list := make([]*DexPosition, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/dex_positions", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *walletClient) ListDexTrades(ctx context.Context, addr Address, params Query) ([]*DexTrade, error) {
	list := make([]*DexTrade, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/dex_trades", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *walletClient) ListFarmEvents(ctx context.Context, addr Address, params Query) ([]*FarmEvent, error) {
	list := make([]*FarmEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/farm_events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *walletClient) ListFarmPositions(ctx context.Context, addr Address, params Query) ([]*FarmPosition, error) {
	list := make([]*FarmPosition, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/farm_positions", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *walletClient) ListLendingEvents(ctx context.Context, addr Address, params Query) ([]*LendingEvent, error) {
	list := make([]*LendingEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/lend_events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *walletClient) ListLendingPositions(ctx context.Context, addr Address, params Query) ([]*LendingPosition, error) {
	list := make([]*LendingPosition, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/lend_positions", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
