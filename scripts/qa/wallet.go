package main

import (
	"context"
	"fmt"

	"blockwatch.cc/tzpro-go/tzpro"
)

func TestWallet(ctx context.Context, c *tzpro.Client) {
	addr := tzpro.NewAddress("tz1go7f6mEQfT2xX2LuHAqgnRGN6c2zHPf5c") // Main
	ap := tzpro.WithMeta()
	op := tzpro.WithStorage().WithMeta()

	// account
	try("GetAccount", func() {
		if _, err := c.Account.Get(ctx, addr, ap); err != nil {
			panic(err)
		}
	})

	// contracts
	try("GetAccountContracts", func() {
		if _, err := c.Account.ListContracts(ctx, addr, ap); err != nil {
			panic(err)
		}
	})

	// ops
	try("GetAccountOps", func() {
		if ops, err := c.Account.ListOps(ctx, addr, op); err != nil || len(ops) == 0 {
			panic(fmt.Errorf("len=%d %v", len(ops), err))
		}
	})

	// account table
	try("Account query", func() {
		aq := c.Account.NewQuery().WithLimit(2).Desc()
		if acc, err := aq.Run(ctx); err != nil {
			panic(err)
		} else if acc.Len() == 0 {
			panic(fmt.Errorf("acc len=%d", acc.Len()))
		}
	})

	// metadata
	try("ListMetadata", func() {
		if _, err := c.Metadata.List(ctx); err != nil {
			panic(err)
		}
	})
	try("GetWalletMetadata", func() {
		if _, err := c.Metadata.GetWallet(ctx, addr); err != nil {
			panic(err)
		}
	})
	try("GetAssetMetadata", func() {
		addr := tzpro.NewToken("KT1XnTn74bUtxHfDtBmm2bGZAQfhPbvKWR8o_0")
		if _, err := c.Metadata.GetAsset(ctx, addr); err != nil {
			panic(err)
		}
	})
	try("Describe", func() {
		addr := tzpro.NewAddress("KT1XnTn74bUtxHfDtBmm2bGZAQfhPbvKWR8o")
		if _, err := c.Metadata.DescribeAddress(ctx, addr); err != nil {
			panic(err)
		}
	})
	try("GetAllMetadataSchemas", func() {
		if _, err := c.Metadata.GetSchemas(ctx); err != nil {
			panic(err)
		}
	})
	try("GetMetadataSchema", func() {
		if _, err := c.Metadata.GetSchema(ctx, "asset"); err != nil {
			panic(err)
		}
	})
}
