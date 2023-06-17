package main

import (
	"context"

	"blockwatch.cc/tzpro-go/tzpro"
)

func TestNft(ctx context.Context, c *tzpro.Client) {
	// dex
	try("ListNftMarkets", func() {
		if _, err := c.ListNftMarkets(ctx, tzpro.NewNftMarketParams()); err != nil {
			panic(err)
		}
	})

	// events
	try("ListNftEvents", func() {
		if _, err := c.ListNftEvents(ctx, tzpro.NewNftEventParams()); err != nil {
			panic(err)
		}
	})

	// positions
	try("ListNftPositions", func() {
		if _, err := c.ListNftPositions(ctx, tzpro.NewNftPositionParams()); err != nil {
			panic(err)
		}
	})

	// trades
	try("ListNftTrades", func() {
		if _, err := c.ListNftTrades(ctx, tzpro.NewNftTradeParams()); err != nil {
			panic(err)
		}
	})
}
