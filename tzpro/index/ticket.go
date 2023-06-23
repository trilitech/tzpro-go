// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

type TicketUpdate struct {
	// Id      TicketUpdateID `json:"row_id"`
	Ticketer Address `json:"ticketer"`
	Type     Prim    `json:"type"`
	Content  Prim    `json:"content"`
	Account  Address `json:"account"`
	Amount   Z       `json:"amount"`
}
