// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"context"
	"fmt"

	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
)

type Event struct {
	// table API only
	RowId     uint64 `json:"row_id"`
	AccountId uint64 `json:"account_id"`
	Height    int64  `json:"height"`
	OpId      uint64 `json:"op_id"`

	// table and explorer API
	Contract tezos.Address  `json:"contract"`
	Type     micheline.Prim `json:"type"        tzpro:"hex"`
	Payload  micheline.Prim `json:"payload"     tzpro:"hex"`
	Tag      string         `json:"tag"`
	TypeHash string         `json:"type_hash"`
}

type EventList struct {
	Rows    []*Event
	columns []string
}

func (l EventList) Len() int {
	return len(l.Rows)
}

func (l EventList) Cursor() uint64 {
	if len(l.Rows) == 0 {
		return 0
	}
	return l.Rows[len(l.Rows)-1].RowId
}

func (l *EventList) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, null) {
		return nil
	}
	if data[0] != '[' {
		return fmt.Errorf("EventList: expected JSON array")
	}
	return DecodeSlice(data, l.columns, &l.Rows)
}

type EventQuery struct {
	tableQuery
}

func (c *Client) NewEventQuery() EventQuery {
	return EventQuery{c.newTableQuery("event", &Event{})}
}

func (q EventQuery) Run(ctx context.Context) (*EventList, error) {
	result := &EventList{
		columns: q.Columns,
	}
	if err := q.client.QueryTable(ctx, &q.tableQuery, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) QueryEvents(ctx context.Context, filter FilterList, cols []string) (*EventList, error) {
	q := c.NewEventQuery()
	if len(cols) > 0 {
		q.Columns = cols
	}
	if len(filter) > 0 {
		q.Filter = filter
	}
	return q.Run(ctx)
}
