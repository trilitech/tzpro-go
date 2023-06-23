// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package market

import (
	"context"

	"blockwatch.cc/tzpro-go/internal/client"
)

type MarketAPI interface {
	GetTicker(context.Context, string, string) (*Ticker, error)
	ListTickers(context.Context) ([]Ticker, error)
	ListCandles(context.Context, CandleParams) (CandleList, error)
}

func NewMarketAPI(c *client.Client) MarketAPI {
	return &marketClient{client: c}
}

type marketClient struct {
	client *client.Client
}
