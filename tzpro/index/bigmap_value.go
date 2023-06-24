// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"blockwatch.cc/tzpro-go/internal/client"
	"blockwatch.cc/tzpro-go/internal/util"
)

type BigmapValue struct {
	RowId     uint64      `json:"row_id"`
	BigmapId  int64       `json:"bigmap_id"`
	KeyId     uint64      `json:"key_id"`
	Hash      ExprHash    `json:"hash"`
	Height    int64       `json:"height"`
	Time      time.Time   `json:"time"`
	Meta      *BigmapMeta `json:"meta,omitempty"        tzpro:"-"`
	Key       MultiKey    `json:"key"                   tzpro:"-"`
	Value     any         `json:"value,omitempty"       tzpro:"-"`
	KeyPrim   *Prim       `json:"key_prim,omitempty"    tzpro:"key,hex"`
	ValuePrim *Prim       `json:"value_prim,omitempty"  tzpro:"value,hex"`
}

func (r BigmapValue) AsKey(typ Type) BigmapKey {
	k, _ := NewKey(typ, *r.KeyPrim)
	return k
}

func (r BigmapValue) AsValue(typ Type) Value {
	return NewValue(typ, *r.ValuePrim)
}

type BigmapMeta struct {
	Contract     Address   `json:"contract"`
	BigmapId     int64     `json:"bigmap_id"`
	UpdateTime   time.Time `json:"time"`
	UpdateHeight int64     `json:"height"`
	UpdateOp     OpHash    `json:"op"`
	Sender       Address   `json:"sender"`
	Source       Address   `json:"source"`
}

func (v BigmapValue) Has(path string) bool {
	return util.HasPath(v.Value, path)
}

func (v BigmapValue) GetString(path string) (string, bool) {
	return util.GetPathString(v.Value, path)
}

func (v BigmapValue) GetInt64(path string) (int64, bool) {
	return util.GetPathInt64(v.Value, path)
}

func (v BigmapValue) GetBig(path string) (*big.Int, bool) {
	return util.GetPathBig(v.Value, path)
}

func (v BigmapValue) GetZ(path string) (Z, bool) {
	return util.GetPathZ(v.Value, path)
}

func (v BigmapValue) GetTime(path string) (time.Time, bool) {
	return util.GetPathTime(v.Value, path)
}

func (v BigmapValue) GetAddress(path string) (Address, bool) {
	return util.GetPathAddress(v.Value, path)
}

func (v BigmapValue) GetValue(path string) (interface{}, bool) {
	return util.GetPathValue(v.Value, path)
}

func (v BigmapValue) Walk(path string, fn util.ValueWalkerFunc) error {
	val := v.Value
	if len(path) > 0 {
		var ok bool
		val, ok = util.GetPathValue(val, path)
		if !ok {
			return nil
		}
	}
	return util.WalkValueMap(path, val, fn)
}

func (v BigmapValue) Unmarshal(val any) error {
	buf, _ := json.Marshal(v.Value)
	return json.Unmarshal(buf, val)
}

type BigmapValueList []*BigmapValue

func (l BigmapValueList) Len() int {
	return len(l)
}

func (l BigmapValueList) Cursor() uint64 {
	if len(l) == 0 {
		return 0
	}
	return l[len(l)-1].RowId
}

type BigmapValueQuery = client.TableQuery[*BigmapValue]

func (c contractClient) NewBigmapValueQuery() *BigmapValueQuery {
	return client.NewTableQuery[*BigmapValue](c.client, "bigmap_values")
}

func (c *contractClient) GetBigmapValue(ctx context.Context, id int64, key string, params Query) (*BigmapValue, error) {
	v := &BigmapValue{}
	u := params.WithPath(fmt.Sprintf("/explorer/bigmap/%d/%s", id, key)).Url()
	if err := c.client.Get(ctx, u, nil, v); err != nil {
		return nil, err
	}
	return v, nil
}

func (c *contractClient) ListBigmapValues(ctx context.Context, id int64, params Query) (BigmapValueList, error) {
	vals := make(BigmapValueList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/bigmap/%d/values", id)).Url()
	if err := c.client.Get(ctx, u, nil, &vals); err != nil {
		return nil, err
	}
	return vals, nil
}
