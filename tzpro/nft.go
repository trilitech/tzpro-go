// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type NftMarket struct {
	Id         uint64    `json:"id"`
	Contract   string    `json:"contract"`
	Kind       string    `json:"kind"`
	Creator    string    `json:"creator"`
	Name       string    `json:"name"`
	Entity     string    `json:"entity"`
	Tags       []string  `json:"tags"`
	FirstBlock int64     `json:"first_block"`
	FirstTime  time.Time `json:"first_time"`
}

type NftMarketParams = Params[NftMarket]

func NewNftMarketParams() NftMarketParams {
	return NftMarketParams{
		Query: make(map[string][]string),
	}
}

func (m NftMarket) Address() tezos.Address {
	a, _ := tezos.ParseAddress(m.Contract)
	return a
}

func (c *Client) GetNftMarket(ctx context.Context, addr tezos.Address, params NftMarketParams) (*NftMarket, error) {
	p := &NftMarket{}
	u := params.WithPath(fmt.Sprintf("/v1/nft/%s", addr)).Url()
	if err := c.get(ctx, u, nil, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (c *Client) ListNftMarkets(ctx context.Context, params NftMarketParams) ([]*NftMarket, error) {
	list := make([]*NftMarket, 0)
	u := params.WithPath("/v1/nft").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
