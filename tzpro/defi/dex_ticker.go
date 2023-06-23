// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"
	"fmt"
	"time"
)

type DexTicker struct {
	Pair             string    `json:"pair"`
	Pool             string    `json:"pool"`
	Name             string    `json:"name"`
	Entity           string    `json:"entity"`
	PriceChange      string    `json:"price_change"`
	PriceChangeBps   string    `json:"price_change_bps"`
	AskPrice         string    `json:"ask_price"`
	WeightedAvgPrice string    `json:"weighted_avg_price"`
	LastPrice        string    `json:"last_price"`
	LastQty          string    `json:"last_qty"`
	LastTradeTime    string    `json:"last_trade_time"`
	BaseVolume       string    `json:"base_volume"`
	QuoteVolume      string    `json:"quote_volume"`
	OpenPrice        string    `json:"open_price"`
	HighPrice        string    `json:"high_price"`
	LowPrice         string    `json:"low_price"`
	OpenTime         time.Time `json:"open_time"`
	CloseTime        time.Time `json:"close_time"`
	NumTrades        int       `json:"num_trades"`
	LiquidityUSD     string    `json:"liquidity_usd"`
	PriceUSD         string    `json:"price_usd"`
}

func (c *dexClient) GetTicker(ctx context.Context, addr PoolAddress) (*DexTicker, error) {
	tick := &DexTicker{}
	u := fmt.Sprintf("/v1/dex/%s/ticker", addr)
	if err := c.client.Get(ctx, u, nil, tick); err != nil {
		return nil, err
	}
	return tick, nil
}

func (c *dexClient) ListTickers(ctx context.Context, params Params) ([]*DexTicker, error) {
	list := make([]*DexTicker, 0)
	u := params.WithPath("/v1/dex/tickers").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
