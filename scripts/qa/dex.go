package main

import (
	"context"

	"blockwatch.cc/tzpro-go/tzpro"
)

func TestDex(ctx context.Context, c *tzpro.Client) {
	// dex
	try("ListDexes", func() {
		if _, err := c.ListDexes(ctx, tzpro.NewDexParams()); err != nil {
			panic(err)
		}
	})

	// events
	try("ListDexEvents", func() {
		if _, err := c.ListDexEvents(ctx, tzpro.NewDexEventParams()); err != nil {
			panic(err)
		}
	})

	// positions
	try("ListDexPositions", func() {
		if _, err := c.ListDexPositions(ctx, tzpro.NewDexPositionParams()); err != nil {
			panic(err)
		}
	})

	// tickers
	try("ListDexTickers", func() {
		if _, err := c.ListDexTickers(ctx, tzpro.NewDexTickerParams()); err != nil {
			panic(err)
		}
	})

	// trades
	try("ListDexTrades", func() {
		if _, err := c.ListDexTrades(ctx, tzpro.NewDexTradeParams()); err != nil {
			panic(err)
		}
	})
}
