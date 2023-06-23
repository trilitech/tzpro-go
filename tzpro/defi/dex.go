// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzpro-go/internal/client"
)

type DexAPI interface {
	GetDex(context.Context, PoolAddress) (*Dex, error)
	GetTicker(context.Context, PoolAddress) (*DexTicker, error)
	ListPoolEvents(context.Context, PoolAddress, Params) ([]*DexEvent, error)
	ListPoolTrades(context.Context, PoolAddress, Params) ([]*DexTrade, error)
	ListPoolPositions(context.Context, PoolAddress, Params) ([]*DexPosition, error)

	// firehose
	ListDex(context.Context, Params) ([]*Dex, error)
	ListTickers(context.Context, Params) ([]*DexTicker, error)
	ListEvents(context.Context, Params) ([]*DexEvent, error)
	ListTrades(context.Context, Params) ([]*DexTrade, error)
	ListPositions(context.Context, Params) ([]*DexPosition, error)
}

func NewDexAPI(c *client.Client) DexAPI {
	return &dexClient{client: c}
}

type dexClient struct {
	client *client.Client
}

type Dex struct {
	Id              uint64    `json:"id"`
	Contract        string    `json:"contract"`
	PairId          int       `json:"pair_id"`
	Creator         string    `json:"creator"`
	Name            string    `json:"name"`
	Entity          string    `json:"entity"`
	Pair            string    `json:"pair"`
	NumTokens       int       `json:"num_tokens"`
	TokenA          *Token    `json:"token_a"`
	TokenB          *Token    `json:"token_b"`
	TokenLP         *Token    `json:"token_lp"`
	FirstBlock      int64     `json:"first_block"`
	FirstTime       time.Time `json:"first_time"`
	Tags            []string  `json:"tags"`
	SupplyA         Z         `json:"supply_a"`
	SupplyB         Z         `json:"supply_b"`
	SupplyLP        Z         `json:"supply_lp"`
	LastChangeBlock int64     `json:"last_change_block"`
	LastChangeTime  time.Time `json:"last_change_time"`
	NumTrades       int       `json:"num_trades"`
	FeesBps         float64   `json:"fees_bps,string"`
	Price           float64   `json:"price,string"`
	PriceUSD        float64   `json:"price_usd,string"`
	LiquidityUSD    float64   `json:"liquidity_usd,string"`
}

func (p Dex) Address() PoolAddress {
	a, _ := ParseAddress(p.Contract)
	return NewPoolAddress(a, p.PairId)
}

func (c *dexClient) GetDex(ctx context.Context, addr PoolAddress) (*Dex, error) {
	p := &Dex{}
	u := fmt.Sprintf("/v1/dex/%s", addr)
	if err := c.client.Get(ctx, u, nil, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (c *dexClient) ListDex(ctx context.Context, params Params) ([]*Dex, error) {
	list := make([]*Dex, 0)
	u := params.WithPath("/v1/dex").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
