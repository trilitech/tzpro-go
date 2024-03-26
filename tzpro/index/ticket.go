// Copyright (c) 2020-2024 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"time"

	"blockwatch.cc/tzpro-go/internal/client"
)

type TicketAPI interface {
	NewQuery() *TicketQuery
	NewUpdateQuery() *TicketUpdateQuery
	NewEventQuery() *TicketEventQuery
	NewBalanceQuery() *TicketBalanceQuery
}

func NewTicketAPI(c *client.Client) TicketAPI {
	return &ticketClient{client: c}
}

type ticketClient struct {
	client *client.Client
}

func (c *ticketClient) NewQuery() *TicketQuery {
	return client.NewTableQuery[*Ticket](c.client, "ticket")
}

func (c *ticketClient) NewUpdateQuery() *TicketUpdateQuery {
	return client.NewTableQuery[*TicketUpdate](c.client, "ticket_updates")
}

func (c *ticketClient) NewBalanceQuery() *TicketBalanceQuery {
	return client.NewTableQuery[*TicketBalance](c.client, "ticket_owners")
}

func (c *ticketClient) NewEventQuery() *TicketEventQuery {
	return client.NewTableQuery[*TicketEvent](c.client, "ticket_events")
}

type (
	TicketList        []*Ticket
	TicketUpdateList  []*TicketUpdate
	TicketBalanceList []*TicketBalance
	TicketEventList   []*TicketEvent

	TicketQuery        = client.TableQuery[*Ticket]
	TicketUpdateQuery  = client.TableQuery[*TicketUpdate]
	TicketBalanceQuery = client.TableQuery[*TicketBalance]
	TicketEventQuery   = client.TableQuery[*TicketEvent]
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
