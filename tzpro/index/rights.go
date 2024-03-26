// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"blockwatch.cc/tzpro-go/internal/client"
	"blockwatch.cc/tzpro-go/internal/util"
)

type Right struct {
	Type           RightType `json:"type"`
	Address        Address   `json:"address"`
	Round          int       `json:"round"`
	IsUsed         bool      `json:"is_used"`
	IsLost         bool      `json:"is_lost"`
	IsStolen       bool      `json:"is_stolen"`
	IsMissed       bool      `json:"is_missed"`
	IsSeedRequired bool      `json:"is_seed_required"`
	IsSeedRevealed bool      `json:"is_seed_revealed"`
}

type Rights struct {
	Id        uint64        `json:"id"`
	Cycle     int64         `json:"cycle"`
	Height    int64         `json:"height"`
	AccountId uint64        `json:"account_id"`
	Address   Address       `json:"address"`
	Bake      util.HexBytes `json:"baking_rights"`
	Endorse   util.HexBytes `json:"endorsing_rights"`
	Baked     util.HexBytes `json:"blocks_baked"`
	Endorsed  util.HexBytes `json:"blocks_endorsed"`
	Seed      util.HexBytes `json:"seeds_required"`
	Seeded    util.HexBytes `json:"seeds_revealed"`
}

func isSet(buf []byte, i int) bool {
	if i < 0 || i >= len(buf)*8 {
		return false
	}
	return (buf[i>>3] & byte(1<<uint(i&0x7))) > 0
}

func (r Rights) Pos(height int64) int {
	return int(height - r.Height)
}

func (r Rights) IsUsed(pos int) bool {
	return isSet(r.Bake, pos) && isSet(r.Baked, pos) || isSet(r.Endorse, pos) && isSet(r.Endorsed, pos)
}

func (r Rights) IsLost(pos int) bool {
	return isSet(r.Bake, pos) && !isSet(r.Baked, pos)
}

func (r Rights) IsStolen(pos int) bool {
	return !isSet(r.Bake, pos) && isSet(r.Baked, pos)
}

func (r Rights) IsMissed(pos int) bool {
	return isSet(r.Endorse, pos) && !isSet(r.Endorsed, pos)
}

func (r Rights) IsSeedRequired(pos int) bool {
	return isSet(r.Seed, pos)
}

func (r Rights) IsSeedRevealed(pos int) bool {
	return isSet(r.Seeded, pos)
}

func (r Rights) RightAt(height int64, typ RightType) (Right, bool) {
	pos := r.Pos(height)
	if typ == RightTypeBaking && (isSet(r.Bake, pos) || isSet(r.Baked, pos)) {
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
	if typ == RightTypeEndorsing && isSet(r.Endorse, pos) {
		return Right{
			Type:     typ,
			Address:  r.Address,
			IsUsed:   isSet(r.Endorse, pos) && isSet(r.Endorsed, pos),
			IsMissed: isSet(r.Endorse, pos) && !isSet(r.Endorsed, pos),
		}, true
	}
	return Right{}, false
}

type RightsList []*Rights

func (l RightsList) Len() int {
	return len(l)
}

func (l RightsList) Cursor() uint64 {
	if len(l) == 0 {
		return 0
	}
	return l[len(l)-1].Id
}

type RightsQuery = client.TableQuery[*Rights]

func (c bakerClient) NewRightsQuery() *RightsQuery {
	return client.NewTableQuery[*Rights](c.client, "rights")
}
