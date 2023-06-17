// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"context"
	"fmt"

	"blockwatch.cc/tzgo/tezos"
)

type Right struct {
	Type           tezos.RightType `json:"type"`
	Address        tezos.Address   `json:"address"`
	Round          int             `json:"round"`
	IsUsed         bool            `json:"is_used"`
	IsLost         bool            `json:"is_lost"`
	IsStolen       bool            `json:"is_stolen"`
	IsMissed       bool            `json:"is_missed"`
	IsSeedRequired bool            `json:"is_seed_required"`
	IsSeedRevealed bool            `json:"is_seed_revealed"`
}

type CycleRights struct {
	RowId     uint64         `json:"row_id"`
	Cycle     int64          `json:"cycle"`
	Height    int64          `json:"height"`
	AccountId uint64         `json:"account_id"`
	Address   tezos.Address  `json:"address"`
	Bake      tezos.HexBytes `json:"baking_rights"`
	Endorse   tezos.HexBytes `json:"endorsing_rights"`
	Baked     tezos.HexBytes `json:"blocks_baked"`
	Endorsed  tezos.HexBytes `json:"blocks_endorsed"`
	Seed      tezos.HexBytes `json:"seeds_required"`
	Seeded    tezos.HexBytes `json:"seeds_revealed"`
}

func isSet(buf []byte, i int) bool {
	if i < 0 || i >= len(buf)*8 {
		return false
	}
	return (buf[i>>3] & byte(1<<uint(i&0x7))) > 0
}

func (r CycleRights) Pos(height int64) int {
	return int(height - r.Height)
}

func (r CycleRights) IsUsed(pos int) bool {
	return isSet(r.Bake, pos) && isSet(r.Baked, pos) || isSet(r.Endorse, pos) && isSet(r.Endorsed, pos)
}

func (r CycleRights) IsLost(pos int) bool {
	return isSet(r.Bake, pos) && !isSet(r.Baked, pos)
}

func (r CycleRights) IsStolen(pos int) bool {
	return !isSet(r.Bake, pos) && isSet(r.Baked, pos)
}

func (r CycleRights) IsMissed(pos int) bool {
	return isSet(r.Endorse, pos) && !isSet(r.Endorsed, pos)
}

func (r CycleRights) IsSeedRequired(pos int) bool {
	return isSet(r.Seed, pos)
}

func (r CycleRights) IsSeedRevealed(pos int) bool {
	return isSet(r.Seeded, pos)
}

func (r CycleRights) RightAt(height int64, typ tezos.RightType) (Right, bool) {
	pos := r.Pos(height)
	if typ == tezos.RightTypeBaking && (isSet(r.Bake, pos) || isSet(r.Baked, pos)) {
		return Right{
			Type:           typ,
			Address:        r.Address,
			IsUsed:         isSet(r.Bake, pos) && isSet(r.Baked, pos),
			IsLost:         isSet(r.Bake, pos) && !isSet(r.Baked, pos),
			IsStolen:       !isSet(r.Bake, pos) && isSet(r.Baked, pos),
			IsSeedRequired: isSet(r.Seed, pos),
			IsSeedRevealed: isSet(r.Seeded, pos),
		}, true
	}
	if typ == tezos.RightTypeEndorsing && isSet(r.Endorse, pos) {
		return Right{
			Type:     typ,
			Address:  r.Address,
			IsUsed:   isSet(r.Endorse, pos) && isSet(r.Endorsed, pos),
			IsMissed: isSet(r.Endorse, pos) && !isSet(r.Endorsed, pos),
		}, true
	}
	return Right{}, false
}

type CycleRightsList struct {
	Rows    []*CycleRights
	columns []string
}

func (l CycleRightsList) Len() int {
	return len(l.Rows)
}

func (l CycleRightsList) Cursor() uint64 {
	if len(l.Rows) == 0 {
		return 0
	}
	return l.Rows[len(l.Rows)-1].RowId
}

func (l *CycleRightsList) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, null) {
		return nil
	}
	if data[0] != '[' {
		return fmt.Errorf("CycleRightsList: expected JSON array")
	}
	return DecodeSlice(data, l.columns, &l.Rows)
}

type CycleRightsQuery struct {
	tableQuery
}

func (c *Client) NewCycleRightsQuery() CycleRightsQuery {
	return CycleRightsQuery{c.newTableQuery("rights", &CycleRights{})}
}

func (q CycleRightsQuery) Run(ctx context.Context) (*CycleRightsList, error) {
	result := &CycleRightsList{
		columns: q.Columns,
	}
	if err := q.client.QueryTable(ctx, &q.tableQuery, result); err != nil {
		return nil, err
	}
	return result, nil
}
