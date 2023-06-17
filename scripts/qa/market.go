package main

import (
	"context"
	"time"

	"blockwatch.cc/tzpro-go/tzpro"
)

func TestMarket(ctx context.Context, c *tzpro.Client) {
	try("GetTickers", func() {
		if _, err := c.GetTickers(ctx); err != nil {
			panic(err)
		}
	})

	try("GetTicker", func() {
		if _, err := c.GetTicker(ctx, "kraken", "XTZ_USD"); err != nil {
			panic(err)
		}
	})
	try("ListCandles", func() {
		args := tzpro.CandleArgs{
			Market:   "kraken",
			Pair:     "XTZ_USD",
			Collapse: tzpro.Collapse1d,
			From:     time.Now().Add(-168 * time.Hour),
		}
		if _, err := c.ListCandles(ctx, args); err != nil {
			panic(err)
		}
	})
}
