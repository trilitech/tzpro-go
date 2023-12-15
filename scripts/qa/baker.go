package main

import (
	"context"
	"fmt"

	"blockwatch.cc/tzpro-go/tzpro"
)

func TestBaker(ctx context.Context, c *tzpro.Client) {
	addr := tzpro.NewAddress("tz1go7f6mEQfT2xX2LuHAqgnRGN6c2zHPf5c") // main
	bkp := tzpro.WithMeta()
	op := tzpro.WithStorage().WithMeta()

	// baker
	try("GetBaker", func() {
		if _, err := c.Baker.Get(ctx, addr, bkp); err != nil {
			panic(err)
		}
	})

	// list
	try("ListBakers", func() {
		if l, err := c.Baker.List(ctx, bkp); err != nil || len(l) == 0 {
			panic(fmt.Errorf("len=%d %v", len(l), err))
		}
	})

	// votes
	try("ListBakerVotes", func() {
		if ops, err := c.Baker.ListVotes(ctx, addr, op); err != nil {
			panic(fmt.Errorf("len=%d %v", len(ops), err))
		}
	})

	// endorse
	try("ListBakerEndorsements", func() {
		if ops, err := c.Baker.ListEndorsements(ctx, addr, op); err != nil {
			panic(fmt.Errorf("len=%d %v", len(ops), err))
		}
	})

	// deleg
	try("ListBakerDelegations", func() {
		if ops, err := c.Baker.ListDelegations(ctx, addr, op); err != nil {
			panic(fmt.Errorf("len=%d %v", len(ops), err))
		}
	})

	// rights
	try("GetBakerRights", func() {
		if _, err := c.Baker.GetRights(ctx, addr, 400, bkp); err != nil {
			panic(err)
		}
	})

	// income
	try("GetBakerIncome", func() {
		if _, err := c.Baker.GetIncome(ctx, addr, 400, bkp); err != nil {
			panic(err)
		}
	})

	// snapshot
	try("GetBakerSnapshot", func() {
		if _, err := c.Baker.GetSnapshot(ctx, addr, 400, bkp); err != nil {
			panic(err)
		}
	})

	// rights table
	try("Rights query", func() {
		rq := c.Baker.NewRightsQuery().WithLimit(2).Desc()
		if _, err := rq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// snapshot table
	try("Snapshot query", func() {
		sq := c.Baker.NewStakeSnapshotQuery().WithLimit(2).Desc()
		if _, err := sq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// -----------------------------------------------------------------
	// Gov
	//
	electionId := 11 // main
	try("GetElection", func() {
		if _, err := c.Explorer.GetElection(ctx, electionId); err != nil {
			panic(err)
		}
	})

	try("ListVoters", func() {
		if _, err := c.Explorer.ListVoters(ctx, electionId, 1); err != nil {
			panic(err)
		}
	})

	try("ListBallots", func() {
		if _, err := c.Explorer.ListBallots(ctx, electionId, 1); err != nil {
			panic(err)
		}
	})
}
