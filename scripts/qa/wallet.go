package main

import (
	"context"
	"fmt"

	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/tzpro"
)

func TestWallet(ctx context.Context, c *tzpro.Client) {
	addr := tezos.MustParseAddress("tz1go7f6mEQfT2xX2LuHAqgnRGN6c2zHPf5c") // Main
	ap := tzpro.NewAccountParams().WithMeta()
	op := tzpro.NewOpParams().WithStorage().WithMeta()

	// account
	try("GetAccount", func() {
		if _, err := c.GetAccount(ctx, addr, ap); err != nil {
			panic(err)
		}
	})

	// contracts
	try("GetAccountContracts", func() {
		if _, err := c.GetAccountContracts(ctx, addr, ap); err != nil {
			panic(err)
		}
	})

	// ops
	try("GetAccountOps", func() {
		if ops, err := c.GetAccountOps(ctx, addr, op); err != nil || len(ops) == 0 {
			panic(fmt.Errorf("len=%d %v", len(ops), err))
		}
	})

	// account table
	try("Account query", func() {
		aq := c.NewAccountQuery()
		aq.WithLimit(2).WithDesc()
		if acc, err := aq.Run(ctx); err != nil {
			panic(err)
		} else if acc.Len() == 0 {
			panic(fmt.Errorf("acc len=%d", acc.Len()))
		} else {
			fmt.Printf("ACC: %#v\n", acc.Rows[0])
		}
	})

	// metadata
	try("ListMetadata", func() {
		if _, err := c.ListMetadata(ctx); err != nil {
			panic(err)
		}
	})
	try("GetWalletMetadata", func() {
		if _, err := c.GetWalletMetadata(ctx, addr); err != nil {
			panic(err)
		}
	})
	try("GetAssetMetadata", func() {
		addr := tezos.MustParseToken("KT1XnTn74bUtxHfDtBmm2bGZAQfhPbvKWR8o_0")
		if _, err := c.GetAssetMetadata(ctx, addr); err != nil {
			panic(err)
		}
	})
	try("Describe", func() {
		addr := tezos.MustParseAddress("KT1XnTn74bUtxHfDtBmm2bGZAQfhPbvKWR8o")
		if _, err := c.DescribeAddress(ctx, addr); err != nil {
			panic(err)
		}
	})
	try("GetAllMetadataSchemas", func() {
		if _, err := c.GetAllMetadataSchemas(ctx); err != nil {
			panic(err)
		}
	})
	try("GetMetadataSchema", func() {
		if _, err := c.GetMetadataSchema(ctx, "asset"); err != nil {
			panic(err)
		}
	})
}
