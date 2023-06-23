package main

import (
	"context"

	"blockwatch.cc/tzpro-go/tzpro"
)

func TestNft(ctx context.Context, c *tzpro.Client) {
	p := tzpro.NewParams()
	// dex
	try("ListNftMarkets", func() {
		if _, err := c.Nft.ListMarkets(ctx, p); err != nil {
			panic(err)
		}
	})

	// events
	try("ListNftEvents", func() {
		if _, err := c.Nft.ListEvents(ctx, p); err != nil {
			panic(err)
		}
	})

	// positions
	try("ListNftPositions", func() {
		if _, err := c.Nft.ListPositions(ctx, p); err != nil {
			panic(err)
		}
	})

	// trades
	try("ListNftTrades", func() {
		if _, err := c.Nft.ListTrades(ctx, p); err != nil {
			panic(err)
		}
	})
}
