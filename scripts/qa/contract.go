package main

import (
	"context"
	"fmt"

	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/tzpro"
)

func TestContract(ctx context.Context, c *tzpro.Client) {
	cp := tzpro.NewParams().WithMeta()
	addr := tezos.MustParseAddress("KT1RJ6PbjHpwc3M5rw5s2Nbmefwbuwbdxton") // main

	// contract
	try("GetContract", func() {
		if _, err := c.Contract.Get(ctx, addr, cp); err != nil {
			panic(err)
		}
	})

	// script
	try("GetContractScript", func() {
		if _, err := c.Contract.GetScript(ctx, addr, cp); err != nil {
			panic(err)
		}
	})

	// storage
	try("GetContractStorage", func() {
		if _, err := c.Contract.GetStorage(ctx, addr, cp); err != nil {
			panic(err)
		}
	})

	// calls
	try("GetContractCalls", func() {
		if _, err := c.Contract.ListCalls(ctx, addr, cp); err != nil {
			panic(err)
		}
	})

	try("Contract query", func() {
		ccq := c.Contract.NewQuery().WithLimit(2).WithDesc()
		if _, err := ccq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// --------------------------------------------------------------
	// Bigmap
	//

	// allocs (find a bigmap with >0 keys)
	var bmid int64 = 511 // HEN ledger
	if _, err := c.Contract.GetBigmap(ctx, bmid, cp); err != nil {
		panic(fmt.Errorf("GetBigmap: %v", err))
	}

	// keys
	try("ListBigmapValues", func() {
		if v, err := c.Contract.ListBigmapValues(ctx, bmid, cp); err != nil {
			panic(err)
		} else {
			if _, err := c.Contract.ListBigmapKeyUpdates(ctx, bmid, v[0].Hash.String(), cp); err != nil {
				panic(fmt.Errorf("ListBigmapKeyUpdates: %v", err))
			}
			// value
			if _, err := c.Contract.GetBigmapValue(ctx, bmid, v[0].Hash.String(), cp); err != nil {
				panic(fmt.Errorf("GetBigmapValue: %v", err))
			}
		}
	})

	// list updates
	try("ListBigmapUpdates", func() {
		if _, err := c.Contract.ListBigmapUpdates(ctx, bmid, cp); err != nil {
			panic(err)
		}
	})

	// bigmap table
	try("Bigmap query", func() {
		bmq := c.Contract.NewBigmapQuery().WithLimit(2).WithDesc()
		if _, err := bmq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// bigmap update table
	try("Bigmap update query", func() {
		bmuq := c.Contract.NewBigmapUpdateQuery().
			WithLimit(2).
			WithDesc().
			WithEqual("bigmap_id", bmid)
		if _, err := bmuq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// bigmap value table
	try("Bigmap value query", func() {
		bmvq := c.Contract.NewBigmapValueQuery().
			WithLimit(2).
			WithDesc().
			WithEqual("bigmap_id", bmid)
		if _, err := bmvq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// -----------------------------------------------------------------
	// Constant
	//
	try("Constant query", func() {
		coq := c.Contract.NewConstantQuery().WithLimit(2).WithDesc()
		if _, err := coq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// -----------------------------------------------------------------
	// Events
	//
	try("Event query", func() {
		coq := c.Contract.NewEventQuery().WithLimit(2).WithDesc()
		if _, err := coq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// -----------------------------------------------------------------
	// Tickets
	//
	// try("Ticket query", func() {
	//     coq := c.NewTicketQuery()
	//     coq.WithLimit(2).WithDesc()
	//     if _, err := coq.Run(ctx); err != nil {
	//         panic(err)
	//     }
	// })

}
