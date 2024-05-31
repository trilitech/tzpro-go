package main

import (
	"context"

	"github.com/trilitech/tzpro-go/tzpro"
)

func TestLend(ctx context.Context, c *tzpro.Client) {
	p := tzpro.NewQuery()
	// dex
	try("ListLendings", func() {
		if _, err := c.Lend.ListPools(ctx, p); err != nil {
			panic(err)
		}
	})

	// events
	try("ListLendingEvents", func() {
		if _, err := c.Lend.ListEvents(ctx, p); err != nil {
			panic(err)
		}
	})

	// positions
	try("ListLendingPositions", func() {
		if _, err := c.Lend.ListPositions(ctx, p); err != nil {
			panic(err)
		}
	})
}
