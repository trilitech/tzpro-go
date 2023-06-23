package main

import (
	"context"
	"fmt"

	"blockwatch.cc/tzpro-go/tzpro"
	"blockwatch.cc/tzpro-go/tzpro/index"
)

func TestCommon(ctx context.Context, c *tzpro.Client) (tip *index.Tip) {
	// fetch status
	try("Status", func() {
		stat, err := c.Explorer.GetStatus(ctx)
		if err != nil {
			panic(err)
		}
		if stat.Status != "synced" {
			panic(fmt.Errorf("Status is %s", stat.Status))
		}
	})

	// tip
	try("Tip", func() {
		if t, err := c.Explorer.GetTip(ctx); err != nil {
			panic(err)
		} else {
			tip = t
		}
	})

	// protocols
	try("ListProtocols", func() {
		if p, err := c.Explorer.ListProtocols(ctx); err != nil || len(p) == 0 {
			panic(fmt.Errorf("len=%d %v", len(p), err))
		}
	})

	// config
	try("GetConfig", func() {
		if _, err := c.Explorer.GetConfigHead(ctx); err != nil {
			panic(err)
		}
	})

	// config from height
	try("GetConfigHeight", func() {
		if _, err := c.Explorer.GetConfigHeight(ctx, tip.Height); err != nil {
			panic(err)
		}
	})

	// Chain
	try("Chain query", func() {
		chq := c.Explorer.NewChainQuery().WithLimit(2).WithDesc()
		if _, err := chq.Run(ctx); err != nil {
			panic(err)
		}
	})

	return
}
