// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
)

type Constant struct {
	RowId       uint64         `json:"row_id"`
	Address     tezos.ExprHash `json:"address"`
	CreatorId   uint64         `json:"creator_id"`
	Creator     tezos.Address  `json:"creator"`
	Height      int64          `json:"height"`
	Time        time.Time      `json:"time"`
	StorageSize int64          `json:"storage_size"`
	Value       micheline.Prim `json:"value"          tzpro:"hex"`
	Features    StringList     `json:"features"`
}

type ConstantList struct {
	Rows    []*Constant
	columns []string
}

func (l ConstantList) Len() int {
	return len(l.Rows)
}

func (l ConstantList) Cursor() uint64 {
	if len(l.Rows) == 0 {
		return 0
	}
	return l.Rows[len(l.Rows)-1].RowId
}

func (l *ConstantList) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, null) {
		return nil
	}
	if data[0] != '[' {
		return fmt.Errorf("ConstantList: expected JSON array")
	}
	return DecodeSlice(data, l.columns, &l.Rows)
}

type ConstantParams = Params[Constant]

func NewConstantParams() ConstantParams {
	return ConstantParams{
		Query: make(map[string][]string),
	}
}

type ConstantQuery struct {
	tableQuery
}

func (c *Client) NewConstantQuery() ConstantQuery {
	return ConstantQuery{c.newTableQuery("constant", &Constant{})}
}

func (q ConstantQuery) Run(ctx context.Context) (*ConstantList, error) {
	result := &ConstantList{
		columns: q.Columns,
	}
	if err := q.client.QueryTable(ctx, &q.tableQuery, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetConstant(ctx context.Context, addr tezos.ExprHash, params ConstantParams) (*Constant, error) {
	cc := &Constant{}
	u := params.WithPath(fmt.Sprintf("/explorer/constant/%s", addr)).Url()
	if err := c.get(ctx, u, nil, cc); err != nil {
		return nil, err
	}
	return cc, nil
}
