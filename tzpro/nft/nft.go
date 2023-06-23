// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package nft

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzpro-go/internal/client"
)

type NftAPI interface {
	GetMarket(context.Context, Address) (*NftMarket, error)
	ListMarketEvents(context.Context, Address, Params) ([]*NftEvent, error)
	ListMarketPositions(context.Context, Address, Params) ([]*NftPosition, error)
	ListMarketTrades(context.Context, Address, Params) ([]*NftTrade, error)

	// firehose
	ListMarkets(context.Context, Params) ([]*NftMarket, error)
	ListEvents(context.Context, Params) ([]*NftEvent, error)
	ListPositions(context.Context, Params) ([]*NftPosition, error)
	ListTrades(context.Context, Params) ([]*NftTrade, error)
}

func NewNftAPI(c *client.Client) NftAPI {
	return &nftClient{client: c}
}

type nftClient struct {
	client *client.Client
}

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

func (m NftMarket) Address() Address {
	a, _ := ParseAddress(m.Contract)
	return a
}

func (c *nftClient) GetMarket(ctx context.Context, addr Address) (*NftMarket, error) {
	p := &NftMarket{}
	u := fmt.Sprintf("/v1/nft/%s", addr)
	if err := c.client.Get(ctx, u, nil, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (c *nftClient) ListMarkets(ctx context.Context, params Params) ([]*NftMarket, error) {
	list := make([]*NftMarket, 0)
	u := params.WithPath("/v1/nft").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
