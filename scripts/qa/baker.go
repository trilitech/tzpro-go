package main

import (
	"context"
	"fmt"

	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/tzpro"
)

func TestBaker(ctx context.Context, c *tzpro.Client) {
	addr := tezos.MustParseAddress("tz1go7f6mEQfT2xX2LuHAqgnRGN6c2zHPf5c") // main
	bkp := tzpro.NewBakerParams().WithMeta()
	op := tzpro.NewOpParams().WithStorage().WithMeta()

	// baker
	try("GetBaker", func() {
		if _, err := c.GetBaker(ctx, addr, bkp); err != nil {
			panic(err)
		}
	})

	// list
	try("ListBakers", func() {
		if l, err := c.ListBakers(ctx, bkp); err != nil || len(l) == 0 {
			panic(fmt.Errorf("len=%d %v", len(l), err))
		}
	})

	// votes
	try("ListBakerVotes", func() {
		if ops, err := c.ListBakerVotes(ctx, addr, op); err != nil {
			panic(fmt.Errorf("len=%d %v", len(ops), err))
		}
	})

	// endorse
	try("ListBakerEndorsements", func() {
		if ops, err := c.ListBakerEndorsements(ctx, addr, op); err != nil {
			panic(fmt.Errorf("len=%d %v", len(ops), err))
		}
	})

	// deleg
	try("ListBakerDelegations", func() {
		if ops, err := c.ListBakerDelegations(ctx, addr, op); err != nil {
			panic(fmt.Errorf("len=%d %v", len(ops), err))
		}
	})

	// rights
	try("ListBakerRights", func() {
		if _, err := c.ListBakerRights(ctx, addr, 400, bkp); err != nil {
			panic(err)
		}
	})

	// income
	try("GetBakerIncome", func() {
		if _, err := c.GetBakerIncome(ctx, addr, 400, bkp); err != nil {
			panic(err)
		}
	})

	// snapshot
	try("GetBakerSnapshot", func() {
		if _, err := c.GetBakerSnapshot(ctx, addr, 400, bkp); err != nil {
			panic(err)
		}
	})

	// rights table
	try("Rights query", func() {
		rq := c.NewCycleRightsQuery()
		rq.WithLimit(2).WithDesc()
		if _, err := rq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// snapshot table
	try("Snapshot query", func() {
		sq := c.NewSnapshotQuery()
		sq.WithLimit(2).WithDesc()
		if _, err := sq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// -----------------------------------------------------------------
	// Gov
	//
	electionId := 11 // main
	try("GetElection", func() {
		if _, err := c.GetElection(ctx, electionId); err != nil {
			panic(err)
		}
	})

	try("ListVoters", func() {
		if _, err := c.ListVoters(ctx, electionId, 1); err != nil {
			panic(err)
		}
	})

	try("ListBallots", func() {
		if _, err := c.ListBallots(ctx, electionId, 1); err != nil {
			panic(err)
		}
	})
}
