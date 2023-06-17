package main

import (
	"context"

	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/tzpro"
)

func TestToken(ctx context.Context, c *tzpro.Client) {
	// ledgers
	try("ListLedgers", func() {
		if _, err := c.ListLedgers(ctx, tzpro.NewLedgerParams()); err != nil {
			panic(err)
		}
	})

	// tokens
	try("ListTokens", func() {
		if _, err := c.ListTokens(ctx, tzpro.NewTokenParams()); err != nil {
			panic(err)
		}
	})

	// balances
	try("ListLedgerBalances", func() {
		addr := tezos.MustParseAddress("KT1RJ6PbjHpwc3M5rw5s2Nbmefwbuwbdxton")
		if _, err := c.ListLedgerBalances(ctx, addr, tzpro.NewTokenBalanceParams()); err != nil {
			panic(err)
		}
	})
	try("ListTokenBalances", func() {
		addr := tezos.MustParseToken("KT1XnTn74bUtxHfDtBmm2bGZAQfhPbvKWR8o_0")
		if _, err := c.ListTokenBalances(ctx, addr, tzpro.NewTokenBalanceParams()); err != nil {
			panic(err)
		}
	})

	// events
	try("ListLedgerEvents", func() {
		addr := tezos.MustParseAddress("KT1RJ6PbjHpwc3M5rw5s2Nbmefwbuwbdxton")
		if _, err := c.ListLedgerEvents(ctx, addr, tzpro.NewTokenEventParams()); err != nil {
			panic(err)
		}
	})
	try("ListTokenEvents", func() {
		if _, err := c.ListTokenEvents(ctx, tzpro.NewTokenEventParams()); err != nil {
			panic(err)
		}
	})
	try("ListTokenIdEvents", func() {
		addr := tezos.MustParseToken("KT1XnTn74bUtxHfDtBmm2bGZAQfhPbvKWR8o_0")
		if _, err := c.ListTokenIdEvents(ctx, addr, tzpro.NewTokenEventParams()); err != nil {
			panic(err)
		}
	})

	// metadata
	try("ListTokenMetadata", func() {
		if _, err := c.ListTokenMetadata(ctx, tzpro.NewTokenMetadataParams()); err != nil {
			panic(err)
		}
	})
	try("GetLedgerMeta", func() {
		addr := tezos.MustParseAddress("KT1RJ6PbjHpwc3M5rw5s2Nbmefwbuwbdxton")
		if _, err := c.GetLedgerMetadata(ctx, addr, tzpro.NewTokenMetadataParams()); err != nil {
			panic(err)
		}
	})
	try("GetTokenMetadata", func() {
		addr := tezos.MustParseToken("KT1XnTn74bUtxHfDtBmm2bGZAQfhPbvKWR8o_0")
		if _, err := c.GetTokenMetadata(ctx, addr, tzpro.NewTokenMetadataParams()); err != nil {
			panic(err)
		}
	})
}
