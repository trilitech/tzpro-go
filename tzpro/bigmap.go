// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"time"

	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
)

type Bigmap struct {
	Contract       tezos.Address     `json:"contract"`
	BigmapId       int64             `json:"bigmap_id"`
	NUpdates       int64             `json:"n_updates"`
	NKeys          int64             `json:"n_keys"`
	AllocateHeight int64             `json:"alloc_height"`
	AllocateBlock  tezos.BlockHash   `json:"alloc_block"`
	AllocateTime   time.Time         `json:"alloc_time"`
	UpdateHeight   int64             `json:"update_height"`
	UpdateBlock    tezos.BlockHash   `json:"update_block"`
	UpdateTime     time.Time         `json:"update_time"`
	DeleteHeight   int64             `json:"delete_height"`
	DeleteBlock    tezos.BlockHash   `json:"delete_block"`
	DeleteTime     time.Time         `json:"delete_time"`
	KeyType        micheline.Typedef `json:"key_type"`
	ValueType      micheline.Typedef `json:"value_type"`
	KeyTypePrim    micheline.Prim    `json:"key_type_prim"`
	ValueTypePrim  micheline.Prim    `json:"value_type_prim"`
}

func (b Bigmap) MakeKeyType() micheline.Type {
	return micheline.NewType(b.KeyTypePrim)
}

func (b Bigmap) MakeValueType() micheline.Type {
	return micheline.NewType(b.ValueTypePrim)
}

type BigmapRow struct {
	RowId        uint64          `json:"row_id"`
	Contract     tezos.Address   `json:"contract"`
	AccountId    uint64          `json:"account_id"`
	BigmapId     int64           `json:"bigmap_id"`
	NUpdates     int64           `json:"n_updates"`
	NKeys        int64           `json:"n_keys"`
	AllocHeight  int64           `json:"alloc_height"`
	AllocTime    time.Time       `json:"alloc_time"`
	AllocBlock   tezos.BlockHash `json:"alloc_block"`
	UpdateHeight int64           `json:"update_height"`
	UpdateTime   time.Time       `json:"update_time"`
	UpdateBlock  tezos.BlockHash `json:"update_block"`
	DeleteHeight int64           `json:"delete_height"`
	DeleteBlock  tezos.BlockHash `json:"delete_block"`
	DeleteTime   time.Time       `json:"delete_time"`
	KeyType      string          `json:"key_type"`
	ValueType    string          `json:"value_type"`
}

func (r BigmapRow) DecodeKeyType() (micheline.Type, error) {
	var t micheline.Type
	buf, err := hex.DecodeString(r.KeyType)
	if err != nil {
		return t, nil
	}
	if len(buf) == 0 {
		return t, io.ErrShortBuffer
	}
	err = t.UnmarshalBinary(buf)
	return t, err
}

func (r BigmapRow) DecodeValueType() (micheline.Type, error) {
	var t micheline.Type
	buf, err := hex.DecodeString(r.ValueType)
	if err != nil {
		return t, nil
	}
	if len(buf) == 0 {
		return t, io.ErrShortBuffer
	}
	err = t.UnmarshalBinary(buf)
	return t, err
}

type BigmapRowList struct {
	Rows    []*BigmapRow
	columns []string
}

func (l BigmapRowList) Len() int {
	return len(l.Rows)
}

func (l BigmapRowList) Cursor() uint64 {
	if len(l.Rows) == 0 {
		return 0
	}
	return l.Rows[len(l.Rows)-1].RowId
}

func (l *BigmapRowList) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, null) {
		return nil
	}
	if data[0] != '[' {
		return fmt.Errorf("BigmapRowList: expected JSON array")
	}
	return DecodeSlice(data, l.columns, &l.Rows)
}

type BigmapQuery struct {
	tableQuery
}

func (c *Client) NewBigmapQuery() BigmapQuery {
	tinfo, err := GetTypeInfo(&BigmapRow{})
	if err != nil {
		panic(err)
	}
	q := tableQuery{
		client:     c,
		BaseParams: c.base.Clone(),
		Table:      "bigmaps",
		Format:     FormatJSON,
		Limit:      DefaultLimit,
		Order:      OrderAsc,
		Columns:    tinfo.Aliases(),
		Filter:     make(FilterList, 0),
	}
	return BigmapQuery{q}
}

func (q BigmapQuery) Run(ctx context.Context) (*BigmapRowList, error) {
	result := &BigmapRowList{
		columns: q.Columns,
	}
	if err := q.client.QueryTable(ctx, &q.tableQuery, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetBigmap(ctx context.Context, id int64, params ContractParams) (*Bigmap, error) {
	b := &Bigmap{}
	u := params.WithPath(fmt.Sprintf("/explorer/bigmap/%d", id)).Url()
	if err := c.get(ctx, u, nil, b); err != nil {
		return nil, err
	}
	return b, nil
}
