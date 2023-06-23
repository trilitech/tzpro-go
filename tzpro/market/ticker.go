// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package market

import (
	"context"
	"fmt"
	"time"
)

type Ticker struct {
	Pair        string    `json:"pair"`
	Base        string    `json:"base"`
	Quote       string    `json:"quote"`
	Exchange    string    `json:"exchange"`
	Open        float64   `json:"open"`
	High        float64   `json:"high"`
	Low         float64   `json:"low"`
	Last        float64   `json:"last"`
	Change      float64   `json:"change"`
	Vwap        float64   `json:"vwap"`
	NTrades     int64     `json:"n_trades"`
	VolumeBase  float64   `json:"volume_base"`
	VolumeQuote float64   `json:"volume_quote"`
	Time        time.Time `json:"timestamp"`
}

func (c *marketClient) ListTickers(ctx context.Context) ([]Ticker, error) {
	ticks := make([]Ticker, 0)
	if err := c.client.Get(ctx, "/markets/tickers", nil, &ticks); err != nil {
		return nil, err
	}
	return ticks, nil
}

func (c *marketClient) GetTicker(ctx context.Context, market, pair string) (*Ticker, error) {
	var tick Ticker
	u := fmt.Sprintf("/markets/%s/%s/ticker", market, pair)
	if err := c.client.Get(ctx, u, nil, &tick); err != nil {
		return nil, err
	}
	return &tick, nil
}
