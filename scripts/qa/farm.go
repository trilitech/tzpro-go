package main

import (
	"context"

	"github.com/trilitech/tzpro-go/tzpro"
)

func TestFarm(ctx context.Context, c *tzpro.Client) {
	p := tzpro.NewQuery()
	// dex
	try("ListFarms", func() {
		if _, err := c.Farm.ListFarms(ctx, p); err != nil {
			panic(err)
		}
	})

	// events
	try("ListFarmEvents", func() {
		if _, err := c.Farm.ListEvents(ctx, p); err != nil {
			panic(err)
		}
	})

	// positions
	try("ListFarmPositions", func() {
		if _, err := c.Farm.ListPositions(ctx, p); err != nil {
			panic(err)
		}
	})
}
