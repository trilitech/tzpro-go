package main

import (
	"context"

	"blockwatch.cc/tzpro-go/tzpro"
)

func TestDex(ctx context.Context, c *tzpro.Client) {
	p := tzpro.NewParams()
	// dex
	try("ListDexes", func() {
		if _, err := c.Dex.ListDex(ctx, p); err != nil {
			panic(err)
		}
	})

	// events
	try("ListDexEvents", func() {
		if _, err := c.Dex.ListEvents(ctx, p); err != nil {
			panic(err)
		}
	})

	// positions
	try("ListDexPositions", func() {
		if _, err := c.Dex.ListPositions(ctx, p); err != nil {
			panic(err)
		}
	})

	// tickers
	try("ListDexTickers", func() {
		if _, err := c.Dex.ListTickers(ctx, p); err != nil {
			panic(err)
		}
	})

	// trades
	try("ListDexTrades", func() {
		if _, err := c.Dex.ListTrades(ctx, p); err != nil {
			panic(err)
		}
	})
}
