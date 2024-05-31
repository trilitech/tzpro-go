package main

import (
	"context"
	"time"

	"github.com/trilitech/tzpro-go/tzpro"
	"github.com/trilitech/tzpro-go/tzpro/market"
)

func TestMarket(ctx context.Context, c *tzpro.Client) {
	try("GetTickers", func() {
		if _, err := c.Market.ListTickers(ctx); err != nil {
			panic(err)
		}
	})

	try("GetTicker", func() {
		if _, err := c.Market.GetTicker(ctx, "kraken", "XTZ_USD"); err != nil {
			panic(err)
		}
	})
	try("ListCandles", func() {
		args := market.CandleQuery{
			Market:   "kraken",
			Pair:     "XTZ_USD",
			Collapse: tzpro.Collapse1d,
			From:     time.Now().Add(-168 * time.Hour),
		}
		if _, err := c.Market.ListCandles(ctx, args); err != nil {
			panic(err)
		}
	})
}
