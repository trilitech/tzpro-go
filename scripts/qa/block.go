package main

import (
	"context"
	"fmt"

	"blockwatch.cc/tzpro-go/tzpro"
)

func TestBlock(ctx context.Context, c *tzpro.Client, tip *tzpro.Tip) {
	bp := tzpro.NewBlockParams().WithRights().WithMeta()
	op := tzpro.NewOpParams().WithStorage().WithMeta()

	// block
	try("GetBlock", func() {
		if _, err := c.GetBlock(ctx, tip.Hash, bp); err != nil {
			panic(err)
		}
	})

	// block head
	try("GetHead", func() {
		if _, err := c.GetHead(ctx, bp); err != nil {
			panic(err)
		}
	})

	// block height
	try("GetBlockHeight", func() {
		if _, err := c.GetBlockHeight(ctx, tip.Height, bp); err != nil {
			panic(err)
		}
	})

	// block ops
	try("GetBlockOps", func() {
		if ops, err := c.GetBlockOps(ctx, tip.Hash, op); err != nil || len(ops) == 0 {
			panic(fmt.Errorf("len=%d %v", len(ops), err))
		}
	})

	// block table
	try("Block query", func() {
		bq := c.NewBlockQuery()
		bq.WithLimit(2).WithDesc()
		if _, err := bq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// -----------------------------------------------------------------
	// Operations
	//
	try("Op query", func() {
		oq := c.NewOpQuery()
		oq.WithFilter(tzpro.FilterModeEqual, "type", "transaction").
			WithLimit(100).
			WithOrder(tzpro.OrderDesc)
		ores, err := oq.Run(ctx)
		if err != nil {
			panic(err)
		}
		if ores.Len() > 0 {
			if _, err := c.GetOp(ctx, ores.Rows[0].Hash, op); err != nil {
				panic(fmt.Errorf("GetOp: %v", err))
			}
		}
	})
}
