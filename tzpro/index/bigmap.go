// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"
	"fmt"
	"time"

	"github.com/trilitech/tzpro-go/internal/client"
)

type Bigmap struct {
	RowId          uint64    `json:"row_id"`
	Contract       Address   `json:"contract"`
	BigmapId       int64     `json:"bigmap_id"`
	NUpdates       int64     `json:"n_updates"`
	NKeys          int64     `json:"n_keys"`
	AllocateHeight int64     `json:"alloc_height"`
	AllocateBlock  BlockHash `json:"alloc_block"`
	AllocateTime   time.Time `json:"alloc_time"`
	UpdateHeight   int64     `json:"update_height"`
	UpdateBlock    BlockHash `json:"update_block"`
	UpdateTime     time.Time `json:"update_time"`
	DeleteHeight   int64     `json:"delete_height"`
	DeleteBlock    BlockHash `json:"delete_block"`
	DeleteTime     time.Time `json:"delete_time"`
	KeyType        Typedef   `json:"key_type"         tzpro:"-"`
	ValueType      Typedef   `json:"value_type"       tzpro:"-"`
	KeyTypePrim    Prim      `json:"key_type_prim"    tzpro:"key_type,hex"`
	ValueTypePrim  Prim      `json:"value_type_prim"  tzpro:"value_type,hex"`
}

func (r Bigmap) GetKeyTypedef() Typedef {
	if !r.KeyType.IsValid() {
		r.KeyType = r.GetKeyType().Typedef("")
	}
	return r.KeyType
}

func (r Bigmap) GetValueTypedef() Typedef {
	if !r.ValueType.IsValid() {
		r.ValueType = r.GetValueType().Typedef("")
	}
	return r.ValueType
}

func (b Bigmap) GetKeyType() Type {
	return NewType(b.KeyTypePrim)
}

func (b Bigmap) GetValueType() Type {
	return NewType(b.ValueTypePrim)
}

type BigmapQuery = client.TableQuery[*Bigmap]

func (c *contractClient) NewBigmapQuery() *BigmapQuery {
	return client.NewTableQuery[*Bigmap](c.client, "bigmaps")
}

func (c *contractClient) GetBigmap(ctx context.Context, id int64, params Query) (*Bigmap, error) {
	b := &Bigmap{}
	u := params.WithPath(fmt.Sprintf("/explorer/bigmap/%d", id)).Url()
	if err := c.client.Get(ctx, u, nil, b); err != nil {
		return nil, err
	}
	return b, nil
}
