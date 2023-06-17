// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"time"

	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
)

type BigmapValue struct {
	Key       MultiKey        `json:"key"`
	Hash      tezos.ExprHash  `json:"hash"`
	Meta      *BigmapMeta     `json:"meta,omitempty"`
	Value     interface{}     `json:"value,omitempty"`
	Height    int64           `json:"height"`
	Time      time.Time       `json:"time"`
	KeyPrim   *micheline.Prim `json:"key_prim,omitempty"`
	ValuePrim *micheline.Prim `json:"value_prim,omitempty"`
}

type BigmapMeta struct {
	Contract     tezos.Address `json:"contract"`
	BigmapId     int64         `json:"bigmap_id"`
	UpdateTime   time.Time     `json:"time"`
	UpdateHeight int64         `json:"height"`
	UpdateOp     tezos.OpHash  `json:"op"`
	Sender       tezos.Address `json:"sender"`
	Source       tezos.Address `json:"source"`
}

func (v BigmapValue) Has(path string) bool {
	return hasPath(v.Value, path)
}

func (v BigmapValue) GetString(path string) (string, bool) {
	return getPathString(v.Value, path)
}

func (v BigmapValue) GetInt64(path string) (int64, bool) {
	return getPathInt64(v.Value, path)
}

func (v BigmapValue) GetBig(path string) (*big.Int, bool) {
	return getPathBig(v.Value, path)
}

func (v BigmapValue) GetZ(path string) (tezos.Z, bool) {
	return getPathZ(v.Value, path)
}

func (v BigmapValue) GetTime(path string) (time.Time, bool) {
	return getPathTime(v.Value, path)
}

func (v BigmapValue) GetAddress(path string) (tezos.Address, bool) {
	return getPathAddress(v.Value, path)
}

func (v BigmapValue) GetValue(path string) (interface{}, bool) {
	return getPathValue(v.Value, path)
}

func (v BigmapValue) Walk(path string, fn ValueWalkerFunc) error {
	val := v.Value
	if len(path) > 0 {
		var ok bool
		val, ok = getPathValue(val, path)
		if !ok {
			return nil
		}
	}
	return walkValueMap(path, val, fn)
}

func (v BigmapValue) Unmarshal(val interface{}) error {
	buf, _ := json.Marshal(v.Value)
	return json.Unmarshal(buf, val)
}

type BigmapValueRow struct {
	RowId    uint64         `json:"row_id"`
	BigmapId int64          `json:"bigmap_id"`
	Height   int64          `json:"height"`
	Time     time.Time      `json:"time"`
	KeyId    uint64         `json:"key_id"`
	Hash     tezos.ExprHash `json:"hash,omitempty"`
	Key      string         `json:"key,omitempty"`
	Value    string         `json:"value,omitempty"`
}

func (r BigmapValueRow) DecodeKey(typ micheline.Type) (micheline.Key, error) {
	buf, err := hex.DecodeString(r.Key)
	if err != nil {
		return micheline.Key{}, err
	}
	if len(buf) == 0 {
		return micheline.Key{}, io.ErrShortBuffer
	}
	return micheline.DecodeKey(typ, buf)
}

func (r BigmapValueRow) DecodeValue(typ micheline.Type) (micheline.Value, error) {
	v := micheline.NewValue(typ, micheline.Prim{})
	buf, err := hex.DecodeString(r.Value)
	if err != nil {
		return v, err
	}
	if len(buf) == 0 {
		return v, io.ErrShortBuffer
	}
	err = v.Decode(buf)
	return v, err
}

type BigmapValueRowList struct {
	Rows    []*BigmapValueRow
	columns []string
}

func (l BigmapValueRowList) Len() int {
	return len(l.Rows)
}

func (l BigmapValueRowList) Cursor() uint64 {
	if len(l.Rows) == 0 {
		return 0
	}
	return l.Rows[len(l.Rows)-1].RowId
}

func (l *BigmapValueRowList) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, null) {
		return nil
	}
	if data[0] != '[' {
		return fmt.Errorf("BigmapValueRowList: expected JSON array")
	}
	return DecodeSlice(data, l.columns, &l.Rows)
}

type BigmapValueQuery struct {
	tableQuery
}

func (c *Client) NewBigmapValueQuery() BigmapValueQuery {
	return BigmapValueQuery{c.newTableQuery("bigmap_values", &BigmapValueRow{})}
}

func (q BigmapValueQuery) Run(ctx context.Context) (*BigmapValueRowList, error) {
	result := &BigmapValueRowList{
		columns: q.Columns,
	}
	if err := q.client.QueryTable(ctx, &q.tableQuery, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetBigmapValue(ctx context.Context, id int64, key string, params ContractParams) (*BigmapValue, error) {
	v := &BigmapValue{}
	u := params.WithPath(fmt.Sprintf("/explorer/bigmap/%d/%s", id, key)).Url()
	if err := c.get(ctx, u, nil, v); err != nil {
		return nil, err
	}
	return v, nil
}

func (c *Client) ListBigmapValues(ctx context.Context, id int64, params ContractParams) ([]BigmapValue, error) {
	vals := make([]BigmapValue, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/bigmap/%d/values", id)).Url()
	if err := c.get(ctx, u, nil, &vals); err != nil {
		return nil, err
	}
	return vals, nil
}
