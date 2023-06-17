package main

import (
	"context"
	"fmt"

	"blockwatch.cc/tzpro-go/tzpro"
)

func TestCommon(ctx context.Context, c *tzpro.Client) (tip *tzpro.Tip) {
	// fetch status
	try("Status", func() {
		stat, err := c.GetStatus(ctx)
		if err != nil {
			panic(err)
		}
		if stat.Status != "synced" {
			panic(fmt.Errorf("Status is %s", stat.Status))
		}
	})

	// tip
	try("Tip", func() {
		if t, err := c.GetTip(ctx); err != nil {
			panic(err)
		} else {
			tip = t
		}
	})

	// protocols
	try("ListProtocols", func() {
		if p, err := c.ListProtocols(ctx); err != nil || len(p) == 0 {
			panic(fmt.Errorf("len=%d %v", len(p), err))
		}
	})

	// config
	try("GetConfig", func() {
		if _, err := c.GetConfig(ctx); err != nil {
			panic(err)
		}
	})

	// config from height
	try("GetConfigHeight", func() {
		if _, err := c.GetConfigHeight(ctx, tip.Height); err != nil {
			panic(err)
		}
	})

	// Chain
	try("Chain query", func() {
		chq := c.NewChainQuery()
		chq.WithLimit(2).WithDesc()
		if _, err := chq.Run(ctx); err != nil {
			panic(err)
		}
	})

	return
}
