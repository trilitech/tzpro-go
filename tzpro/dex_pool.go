// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

//nolint:staticcheck
type DexPool struct {
	Id              uint64        `json:"id"`
	Contract        tezos.Address `json:"contract"`
	PairId          int           `json:"pair_id"`
	Creator         string        `json:"creator"`
	Name            string        `json:"name"`
	Entity          string        `json:"entity"`
	Pair            string        `json:"pair"`
	NumTokens       int           `json:"num_tokens"`
	TokenA          *Token        `json:"token_a"`
	TokenB          *Token        `json:"token_b"`
	TokenLP         *Token        `json:"token_lp"`
	FirstBlock      int64         `json:"first_block"`
	FirstTime       time.Time     `json:"first_time"`
	Tags            []string      `json:"tags"`
	SupplyA         tezos.Z       `json:"supply_a"`
	SupplyB         tezos.Z       `json:"supply_b"`
	SupplyLP        tezos.Z       `json:"supply_lp"`
	LastChangeBlock int64         `json:"last_change_block"`
	LastChangeTime  time.Time     `json:"last_change_time"`
	NumTrades       int           `json:"num_trades"`
	FeesBps         float64       `json:"fees_bps,string"`
	Price           float64       `json:"price,string"`
	PriceUSD        float64       `json:"price_usd,string"`
	LiquidityUSD    float64       `json:"liquidity_usd,string"`
}

type DexPoolParams = Params[DexPool]

func NewDexPoolParams() DexPoolParams {
	return DexPoolParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) GetDexPool(ctx context.Context, addr tezos.Address, id int, params DexPoolParams) (*DexPool, error) {
	p := &DexPool{}
	u := params.WithPath(fmt.Sprintf("/v1/dex/pools/%s_%d", addr, id)).Url()
	if err := c.get(ctx, u, nil, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (c *Client) ListDexPools(ctx context.Context, params DexPoolParams) ([]*DexPool, error) {
	list := make([]*DexPool, 0)
	u := params.WithPath("/v1/dex/pools").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
