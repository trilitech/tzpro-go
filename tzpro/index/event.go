// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"github.com/trilitech/tzpro-go/internal/client"
)

type Event struct {
	// table API only
	RowId     uint64 `json:"row_id"`
	AccountId uint64 `json:"account_id"`
	Height    int64  `json:"height"`
	OpId      uint64 `json:"op_id"`

	// table and explorer API
	Contract Address `json:"contract"`
	Type     Prim    `json:"type"        tzpro:",hex"`
	Payload  Prim    `json:"payload"     tzpro:",hex"`
	Tag      string  `json:"tag"`
	TypeHash string  `json:"type_hash"`
}

type EventQuery = client.TableQuery[*Event]

func (a contractClient) NewEventQuery() *EventQuery {
	return client.NewTableQuery[*Event](a.client, "event")
}
