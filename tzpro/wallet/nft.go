// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package wallet

import (
	"context"
	"fmt"
)

func (c *walletClient) ListNftEvents(ctx context.Context, addr Address, params Query) ([]*NftEvent, error) {
	list := make([]*NftEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/nft_events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *walletClient) ListNftPositions(ctx context.Context, addr Address, params Query) ([]*NftPosition, error) {
	list := make([]*NftPosition, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/nft_positions", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *walletClient) ListNftTrades(ctx context.Context, addr Address, params Query) ([]*NftTrade, error) {
	list := make([]*NftTrade, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/nft_trades", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
