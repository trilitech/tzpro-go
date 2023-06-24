package main

import (
	"context"
	"fmt"

	"blockwatch.cc/tzpro-go/tzpro"
	"blockwatch.cc/tzpro-go/tzpro/index"
)

func TestBlock(ctx context.Context, c *tzpro.Client, tip *index.Tip) {
	bp := tzpro.WithRights().WithMeta()
	op := bp.Clone()

	// block
	try("GetBlock", func() {
		if _, err := c.Block.GetHash(ctx, tip.Hash, bp); err != nil {
			panic(err)
		}
	})

	// block head
	try("GetHead", func() {
		if _, err := c.Block.GetHead(ctx, bp); err != nil {
			panic(err)
		}
	})

	// block height
	try("GetBlockHeight", func() {
		if _, err := c.Block.GetHeight(ctx, tip.Height, bp); err != nil {
			panic(err)
		}
	})

	// block ops
	try("GetBlockOps", func() {
		if ops, err := c.Block.ListOpsHash(ctx, tip.Hash, op); err != nil || len(ops) == 0 {
			panic(fmt.Errorf("len=%d %v", len(ops), err))
		}
	})

	// block table
	try("Block query", func() {
		bq := c.Block.NewQuery().WithLimit(2).WithDesc()
		if _, err := bq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// -----------------------------------------------------------------
	// Operations
	//
	try("Op query", func() {
		oq := c.Op.NewQuery().
			AndEqual("type", "transaction").
			WithLimit(100).
			WithDesc()
		ores, err := oq.Run(ctx)
		if err != nil {
			panic(err)
		}
		if ores.Len() > 0 {
			if _, err := c.Op.Get(ctx, ores.Rows()[0].Hash, op); err != nil {
				panic(fmt.Errorf("GetOp: %v", err))
			}
		}
	})
}
