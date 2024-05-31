// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package nft

import (
	"context"
	"fmt"
	"time"

	"github.com/trilitech/tzpro-go/internal/client"
)

type NftAPI interface {
	GetMarket(context.Context, Address) (*NftMarket, error)
	ListMarketEvents(context.Context, Address, Query) ([]*NftEvent, error)
	ListMarketPositions(context.Context, Address, Query) ([]*NftPosition, error)
	ListMarketTrades(context.Context, Address, Query) ([]*NftTrade, error)

	// firehose
	ListMarkets(context.Context, Query) ([]*NftMarket, error)
	ListEvents(context.Context, Query) ([]*NftEvent, error)
	ListPositions(context.Context, Query) ([]*NftPosition, error)
	ListTrades(context.Context, Query) ([]*NftTrade, error)
}

func NewNftAPI(c *client.Client) NftAPI {
	return &nftClient{client: c}
}

type nftClient struct {
	client *client.Client
}

type NftMarket struct {
	Id         uint64    `json:"id"`
	Contract   Address   `json:"contract"`
	Kind       string    `json:"kind"`
	Creator    Address   `json:"creator"`
	Name       string    `json:"name"`
	Entity     string    `json:"entity"`
	Tags       []string  `json:"tags"`
	FirstBlock int64     `json:"first_block"`
	FirstTime  time.Time `json:"first_time"`
}

func (c *nftClient) GetMarket(ctx context.Context, addr Address) (*NftMarket, error) {
	p := &NftMarket{}
	u := fmt.Sprintf("/v1/nft/%s", addr)
	if err := c.client.Get(ctx, u, nil, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (c *nftClient) ListMarkets(ctx context.Context, params Query) ([]*NftMarket, error) {
	list := make([]*NftMarket, 0)
	u := params.WithPath("/v1/nft").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
