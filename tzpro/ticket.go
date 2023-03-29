// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
)

type TicketUpdate struct {
	// Id      TicketUpdateID `json:"row_id"`
	Ticketer tezos.Address  `json:"ticketer"`
	Type     micheline.Prim `json:"type"`
	Content  micheline.Prim `json:"content"`
	Account  tezos.Address  `json:"account"`
	Amount   tezos.Z        `json:"amount"`
}
