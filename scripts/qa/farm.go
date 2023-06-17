package main

import (
	"context"

	"blockwatch.cc/tzpro-go/tzpro"
)

func TestFarm(ctx context.Context, c *tzpro.Client) {
	// dex
	try("ListFarms", func() {
		if _, err := c.ListFarms(ctx, tzpro.NewFarmParams()); err != nil {
			panic(err)
		}
	})

	// events
	try("ListFarmEvents", func() {
		if _, err := c.ListFarmEvents(ctx, tzpro.NewFarmEventParams()); err != nil {
			panic(err)
		}
	})

	// positions
	try("ListFarmPositions", func() {
		if _, err := c.ListFarmPositions(ctx, tzpro.NewFarmPositionParams()); err != nil {
			panic(err)
		}
	})
}
