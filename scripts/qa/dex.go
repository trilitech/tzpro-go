package main

import (
	"context"

	"blockwatch.cc/tzpro-go/tzpro"
)

func TestDex(ctx context.Context, c *tzpro.Client) {
	p := tzpro.NewQuery()
	// dex
	try("ListDexes", func() {
		if _, err := c.Dex.ListDex(ctx, p); err != nil {
			panic(err)
		}
	})

	addr := tzpro.NewPoolAddres("KT1J8Hr3BP8bpbfmgGpRPoC9nAMSYtStZG43_0")
	try("GetDex", func() {
		if _, err := c.Dex.GetDex(ctx, addr); err != nil {
			panic(err)
		}
	})

	// events
	try("ListDexEvents", func() {
		if _, err := c.Dex.ListEvents(ctx, p); err != nil {
			panic(err)
		}
	})
	try("ListDexPoolEvents", func() {
		if _, err := c.Dex.ListPoolEvents(ctx, addr, p); err != nil {
			panic(err)
		}
	})

	// positions
	try("ListDexPositions", func() {
		if _, err := c.Dex.ListPositions(ctx, p); err != nil {
			panic(err)
		}
	})
	try("ListDexPoolPositions", func() {
		if _, err := c.Dex.ListPoolPositions(ctx, addr, p); err != nil {
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
	try("ListDexPoolTrades", func() {
		if _, err := c.Dex.ListPoolTrades(ctx, addr, p); err != nil {
			panic(err)
		}
	})
}
