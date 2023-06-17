// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type Snapshot struct {
	RowId        uint64        `json:"row_id"`
	Height       int64         `json:"height"`
	Cycle        int64         `json:"cycle"`
	IsSelected   bool          `json:"is_selected"`
	Timestamp    time.Time     `json:"time"`
	Index        int64         `json:"index"`
	Rolls        int64         `json:"rolls"`
	AccountId    uint64        `json:"account_id"`
	Address      tezos.Address `json:"address"`
	BakerId      uint64        `json:"baker_id"`
	Baker        tezos.Address `json:"baker"`
	IsBaker      bool          `json:"is_baker"`
	IsActive     bool          `json:"is_active"`
	Balance      float64       `json:"balance"`
	Delegated    float64       `json:"delegated"`
	NDelegations int64         `json:"n_delegations"`
	Since        int64         `json:"since"`
	SinceTime    time.Time     `json:"since_time"`
}

type SnapshotList struct {
	Rows    []*Snapshot
	columns []string
}

func (l SnapshotList) Len() int {
	return len(l.Rows)
}

func (l SnapshotList) Cursor() uint64 {
	if len(l.Rows) == 0 {
		return 0
	}
	return l.Rows[len(l.Rows)-1].RowId
}

func (l *SnapshotList) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, null) {
		return nil
	}
	if data[0] != '[' {
		return fmt.Errorf("SnapshotList: expected JSON array")
	}
	return DecodeSlice(data, l.columns, &l.Rows)
}

type SnapshotQuery struct {
	tableQuery
}

func (c *Client) NewSnapshotQuery() SnapshotQuery {
	return SnapshotQuery{c.newTableQuery("snapshot", &Snapshot{})}
}

func (q SnapshotQuery) Run(ctx context.Context) (*SnapshotList, error) {
	result := &SnapshotList{
		columns: q.Columns,
	}
	if err := q.client.QueryTable(ctx, &q.tableQuery, result); err != nil {
		return nil, err
	}
	return result, nil
}
