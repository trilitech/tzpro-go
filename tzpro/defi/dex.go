// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"
	"fmt"
	"time"

	"github.com/trilitech/tzpro-go/internal/client"
)

type DexAPI interface {
	GetDex(context.Context, PoolAddress) (*Dex, error)
	GetTicker(context.Context, PoolAddress) (*DexTicker, error)
	ListPoolEvents(context.Context, PoolAddress, Query) ([]*DexEvent, error)
	ListPoolTrades(context.Context, PoolAddress, Query) ([]*DexTrade, error)
	ListPoolPositions(context.Context, PoolAddress, Query) ([]*DexPosition, error)

	// firehose
	ListDex(context.Context, Query) ([]*Dex, error)
	ListTickers(context.Context, Query) ([]*DexTicker, error)
	ListEvents(context.Context, Query) ([]*DexEvent, error)
	ListTrades(context.Context, Query) ([]*DexTrade, error)
	ListPositions(context.Context, Query) ([]*DexPosition, error)
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
	LastTradeTime   time.Time `json:"last_trade_time"`
	NumTrades       int       `json:"num_trades"`
	FeesBps         float64   `json:"fees_bps,string"`
	Price           float64   `json:"price,string"`
	PriceUSD        float64   `json:"price_usd,string"`
	LiquidityUSD    float64   `json:"liquidity_usd,string"`
	NumCurrentLPs   int       `json:"num_lps"`
	NumTotalLPs     int       `json:"num_total_lps"`
	NumTotalUsers   int       `json:"num_total_users"`
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

func (c *dexClient) ListDex(ctx context.Context, params Query) ([]*Dex, error) {
	list := make([]*Dex, 0)
	u := params.WithPath("/v1/dex").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
