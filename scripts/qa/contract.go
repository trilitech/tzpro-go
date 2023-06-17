package main

import (
	"context"
	"fmt"

	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/tzpro"
)

func TestContract(ctx context.Context, c *tzpro.Client) {
	cp := tzpro.NewContractParams().WithMeta()
	addr := tezos.MustParseAddress("KT1RJ6PbjHpwc3M5rw5s2Nbmefwbuwbdxton") // main

	// contract
	try("GetContract", func() {
		if _, err := c.GetContract(ctx, addr, cp); err != nil {
			panic(err)
		}
	})

	// script
	try("GetContractScript", func() {
		if _, err := c.GetContractScript(ctx, addr, cp); err != nil {
			panic(err)
		}
	})

	// storage
	try("GetContractStorage", func() {
		if _, err := c.GetContractStorage(ctx, addr, cp); err != nil {
			panic(err)
		}
	})

	// calls
	try("GetContractCalls", func() {
		if _, err := c.ListContractCalls(ctx, addr, cp); err != nil {
			panic(err)
		}
	})

	try("Contract query", func() {
		ccq := c.NewContractQuery()
		ccq.WithLimit(2).WithDesc()
		if _, err := ccq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// --------------------------------------------------------------
	// Bigmap
	//

	// allocs (find a bigmap with >0 keys)
	var bmid int64 = 511 // HEN ledger
	if _, err := c.GetBigmap(ctx, bmid, cp); err != nil {
		panic(fmt.Errorf("GetBigmap: %v", err))
	}

	// keys
	try("ListBigmapKeys", func() {
		if k, err := c.ListBigmapKeys(ctx, bmid, cp); err != nil {
			panic(err)
		} else {
			if _, err := c.ListBigmapKeyUpdates(ctx, bmid, k[0].KeyHash.String(), cp); err != nil {
				panic(fmt.Errorf("ListBigmapKeyUpdates: %v", err))
			}
			// value
			if _, err := c.GetBigmapValue(ctx, bmid, k[0].KeyHash.String(), cp); err != nil {
				panic(fmt.Errorf("GetBigmapValue: %v", err))
			}
		}
	})

	// list values
	try("ListBigmapValues", func() {
		if _, err := c.ListBigmapValues(ctx, bmid, cp); err != nil {
			panic(err)
		}
	})

	// list updates
	try("ListBigmapUpdates", func() {
		if _, err := c.ListBigmapUpdates(ctx, bmid, cp); err != nil {
			panic(err)
		}
	})

	// bigmap table
	try("Bigmap query", func() {
		bmq := c.NewBigmapQuery()
		bmq.WithLimit(2).WithDesc()
		if _, err := bmq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// bigmap update table
	try("Bigmap update query", func() {
		bmuq := c.NewBigmapUpdateQuery()
		bmuq.WithLimit(2).WithDesc().WithFilter(tzpro.FilterModeEqual, "bigmap_id", bmid)
		if _, err := bmuq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// bigmap value table
	try("Bigmap value query", func() {
		bmvq := c.NewBigmapValueQuery()
		bmvq.WithLimit(2).WithDesc().WithFilter(tzpro.FilterModeEqual, "bigmap_id", bmid)
		if _, err := bmvq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// -----------------------------------------------------------------
	// Constant
	//
	try("Constant query", func() {
		coq := c.NewConstantQuery()
		coq.WithLimit(2).WithDesc()
		if _, err := coq.Run(ctx); err != nil {
			panic(err)
		}
	})

	// -----------------------------------------------------------------
	// Events
	//
	try("Event query", func() {
		coq := c.NewEventQuery()
		coq.WithLimit(2).WithDesc()
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
