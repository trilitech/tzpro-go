// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzpro-go/internal/client"
)

type BigmapUpdate struct {
	BigmapValue
	Action        DiffAction `json:"action"`
	KeyType       *Typedef   `json:"key_type,omitempty"`
	ValueType     *Typedef   `json:"value_type,omitempty"`
	KeyTypePrim   *Prim      `json:"key_type_prim,omitempty"`
	ValueTypePrim *Prim      `json:"value_type_prim,omitempty"`
	BigmapId      int64      `json:"bigmap_id"`
	SourceId      int64      `json:"source_big_map,omitempty"`
	DestId        int64      `json:"destination_big_map,omitempty"`
}

type BigmapUpdateList []*BigmapUpdate

func (u BigmapUpdate) Event() (ev BigmapEvent) {
	ev.Action = u.Action
	ev.Id = u.BigmapId
	ev.SourceId = u.SourceId
	ev.DestId = u.DestId
	switch u.Action {
	case DiffActionAlloc, DiffActionCopy:
		if u.KeyTypePrim != nil {
			ev.KeyType = *u.KeyTypePrim
		}
		if u.ValueTypePrim != nil {
			ev.ValueType = *u.ValueTypePrim
		}
	case DiffActionUpdate:
		if u.KeyPrim != nil {
			ev.Key = *u.KeyPrim
		}
		if u.ValuePrim != nil {
			ev.Value = *u.ValuePrim
		}
	case DiffActionRemove:
		if u.KeyPrim != nil {
			ev.Key = *u.KeyPrim
		}
	}
	return
}

func (l BigmapUpdateList) Events() (ev BigmapEvents) {
	for _, v := range l {
		ev = append(ev, v.Event())
	}
	return
}

// BigmapUpdateRow is a custom type for table query results which contain
// raw bigmap data
type BigmapUpdateRow struct {
	RowId    uint64     `json:"row_id"`
	BigmapId int64      `json:"bigmap_id"`
	KeyId    uint64     `json:"key_id"`
	Action   DiffAction `json:"action"`
	Height   int64      `json:"height"`
	Time     time.Time  `json:"time"`
	Hash     ExprHash   `json:"hash,omitempty"`
	Key      Prim       `json:"key,omitempty"     tzpro:",hex"`
	Value    Prim       `json:"value,omitempty"   tzpro:",hex"`
}

func (r BigmapUpdateRow) Event() (ev BigmapEvent) {
	ev.Action = r.Action
	ev.Id = r.BigmapId
	switch r.Action {
	case DiffActionCopy:
		ev.SourceId = int64(r.KeyId)
		ev.DestId = r.BigmapId
	case DiffActionAlloc:
		ev.KeyType = r.Key
		ev.ValueType = r.Value
	case DiffActionUpdate:
		ev.Key = r.Key
		ev.Value = r.Value
	case DiffActionRemove:
		ev.Key = r.Key
	}
	return
}

// Alloc/Copy only
func (r BigmapUpdateRow) KeyType() (Type, bool) {
	switch r.Action {
	case DiffActionAlloc, DiffActionCopy:
		return NewType(r.Key), true
	default:
		return Type{}, false
	}
}

// Alloc/Copy only
func (r BigmapUpdateRow) ValueType() (Type, bool) {
	switch r.Action {
	case DiffActionAlloc, DiffActionCopy:
		return NewType(r.Value), true
	default:
		return Type{}, false
	}
}

// Update/Remove only
func (r BigmapUpdateRow) AsKey(typ Type) (BigmapKey, bool) {
	switch r.Action {
	case DiffActionUpdate, DiffActionRemove:
		k, err := NewKey(typ, r.Key)
		return k, err == nil
	default:
		return BigmapKey{}, false
	}
}

// Update only
func (r BigmapUpdateRow) AsValue(typ Type) (Value, bool) {
	if r.Action != DiffActionUpdate {
		return Value{}, false
	}
	return NewValue(typ, r.Value), true
}

type BigmapUpdateQuery = client.TableQuery[*BigmapUpdateRow]

func (c *contractClient) NewBigmapUpdateQuery() *BigmapUpdateQuery {
	return client.NewTableQuery[*BigmapUpdateRow](c.client, "bigmap_updates")
}

func (c *contractClient) ListBigmapUpdates(ctx context.Context, id int64, params Params) (BigmapUpdateList, error) {
	upd := make(BigmapUpdateList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/bigmap/%d/updates", id)).Url()
	if err := c.client.Get(ctx, u, nil, &upd); err != nil {
		return nil, err
	}
	return upd, nil
}

func (c *contractClient) ListBigmapKeyUpdates(ctx context.Context, id int64, key string, params Params) (BigmapUpdateList, error) {
	upd := make(BigmapUpdateList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/bigmap/%d/%s/updates", id, key)).Url()
	if err := c.client.Get(ctx, u, nil, &upd); err != nil {
		return nil, err
	}
	return upd, nil
}
