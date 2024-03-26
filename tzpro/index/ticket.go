// Copyright (c) 2020-2024 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"time"
)

type (
	TicketList        []*Ticket
	TicketUpdateList  []*TicketUpdate
	TicketBalanceList []*TicketBalance
	TicketEventList   []*TicketEvent
)

func (l TicketList) Len() int {
	return len(l)
}

func (l TicketList) Cursor() uint64 {
	if len(l) == 0 {
		return 0
	}
	return l[len(l)-1].Id
}

func (l TicketUpdateList) Len() int {
	return len(l)
}

func (l TicketUpdateList) Cursor() uint64 {
	if len(l) == 0 {
		return 0
	}
	return l[len(l)-1].Id
}

func (l TicketBalanceList) Len() int {
	return len(l)
}

func (l TicketBalanceList) Cursor() uint64 {
	if len(l) == 0 {
		return 0
	}
	return l[len(l)-1].Id
}

func (l TicketEventList) Len() int {
	return len(l)
}

func (l TicketEventList) Cursor() uint64 {
	if len(l) == 0 {
		return 0
	}
	return l[len(l)-1].Id
}

type TicketUpdate struct {
	Id       uint64  `json:"id"`
	Ticketer Address `json:"ticketer"`
	Type     Prim    `json:"type"`
	Content  Prim    `json:"content"`
	Account  Address `json:"account"`
	Amount   Z       `json:"amount"`
}

type Ticket struct {
	Id           uint64    `json:"id"`
	Ticketer     Address   `json:"ticketer"`
	Type         Prim      `json:"type"            tzpro:"type,hex"`
	Content      Prim      `json:"content"         tzpro:"content,hex"`
	Hash         string    `json:"hash"`
	Creator      Address   `json:"creator"`
	FirstBlock   int64     `json:"first_block"`
	FirstTime    time.Time `json:"first_time"`
	LastBlock    int64     `json:"last_block"`
	LastTime     time.Time `json:"last_time"`
	Supply       Z         `json:"total_supply"`
	TotalMint    Z         `json:"total_mint"`
	TotalBurn    Z         `json:"total_burn"`
	NumTransfers int       `json:"num_transfers"`
	NumHolders   int       `json:"num_holders"`
}

type TicketBalance struct {
	Id           uint64    `json:"id"`
	TicketId     uint64    `json:"-"               tzpro:"ticket"`
	Ticketer     Address   `json:"ticketer"        tzpro:"-"`
	Type         Prim      `json:"type"            tzpro:"-"`
	Content      Prim      `json:"content"         tzpro:"-"`
	Hash         string    `json:"hash"            tzpro:"-"`
	Account      Address   `json:"account"`
	Balance      Z         `json:"balance"`
	FirstBlock   int64     `json:"first_block"`
	FirstTime    time.Time `json:"first_time"`
	LastBlock    int64     `json:"last_block"`
	LastTime     time.Time `json:"last_time"`
	NumTransfers int       `json:"num_transfers"`
	NumMints     int       `json:"num_mints"`
	NumBurns     int       `json:"num_burns"`
	VolSent      Z         `json:"vol_sent"`
	VolRecv      Z         `json:"vol_recv"`
	VolMint      Z         `json:"vol_mint"`
	VolBurn      Z         `json:"vol_burn"`
}

type TicketEvent struct {
	Id        uint64    `json:"id"`
	TicketId  uint64    `json:"-"           tzpro:"ticket"`
	Ticketer  Address   `json:"ticketer"    tzpro:"-"`
	Type      Prim      `json:"type"        tzpro:"-"`
	Content   Prim      `json:"content"     tzpro:"-"`
	Hash      string    `json:"hash"        tzpro:"-"`
	EventType string    `json:"event_type"  tzpro:"type"`
	Sender    Address   `json:"sender"`
	Receiver  Address   `json:"receiver"`
	Amount    Z         `json:"amount"`
	Height    int64     `json:"height"`
	Time      time.Time `json:"time"`
	OpId      uint64    `json:"op_id"`
}
