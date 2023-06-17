package main

import (
	"context"

	"blockwatch.cc/tzpro-go/tzpro"
)

func TestLend(ctx context.Context, c *tzpro.Client) {
	// dex
	try("ListLendings", func() {
		if _, err := c.ListLendingPools(ctx, tzpro.NewLendingPoolParams()); err != nil {
			panic(err)
		}
	})

	// events
	try("ListLendingEvents", func() {
		if _, err := c.ListLendingEvents(ctx, tzpro.NewLendingEventParams()); err != nil {
			panic(err)
		}
	})

	// positions
	try("ListLendingPositions", func() {
		if _, err := c.ListLendingPositions(ctx, tzpro.NewLendingPositionParams()); err != nil {
			panic(err)
		}
	})
}
