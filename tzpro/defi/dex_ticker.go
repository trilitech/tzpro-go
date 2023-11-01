// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"
	"fmt"
	"time"
)

type DexTicker struct {
	Pair             string      `json:"pair"`
	Pool             PoolAddress `json:"pool"`
	Name             string      `json:"name"`
	Entity           string      `json:"entity"`
	PriceChange      float64     `json:"price_change,string"`
	PriceChangeBps   float64     `json:"price_change_bps,string"`
	AskPrice         float64     `json:"ask_price,string"`
	WeightedAvgPrice float64     `json:"weighted_avg_price,string"`
	LastPrice        float64     `json:"last_price,string"`
	LastQty          float64     `json:"last_qty,string"`
	LastTradeTime    time.Time   `json:"last_trade_time"`
	BaseVolume       float64     `json:"base_volume,string"`
	QuoteVolume      float64     `json:"quote_volume,string"`
	OpenPrice        float64     `json:"open_price,string"`
	HighPrice        float64     `json:"high_price,string"`
	LowPrice         float64     `json:"low_price,string"`
	OpenTime         time.Time   `json:"open_time"`
	CloseTime        time.Time   `json:"close_time"`
	NumTrades        int         `json:"num_trades"`
	LiquidityUSD     string      `json:"liquidity_usd"`
	PriceUSD         string      `json:"price_usd"`
}

func (c *dexClient) GetTicker(ctx context.Context, addr PoolAddress) (*DexTicker, error) {
	tick := &DexTicker{}
	u := fmt.Sprintf("/v1/dex/%s/ticker", addr)
	if err := c.client.Get(ctx, u, nil, tick); err != nil {
		return nil, err
	}
	return tick, nil
}

func (c *dexClient) ListTickers(ctx context.Context, params Query) ([]*DexTicker, error) {
	list := make([]*DexTicker, 0)
	u := params.WithPath("/v1/dex/tickers").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
