// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"time"

	"github.com/trilitech/tzpro-go/internal/client"
)

type StakeSnapshot struct {
	RowId          uint64    `json:"row_id"`
	Height         int64     `json:"height"`
	Cycle          int64     `json:"cycle"`
	Timestamp      time.Time `json:"time"`
	Index          int64     `json:"index"`
	AccountId      uint64    `json:"account_id"`
	Address        Address   `json:"address"`
	BakerId        uint64    `json:"baker_id"`
	Baker          Address   `json:"baker"`
	IsBaker        bool      `json:"is_baker"`
	IsActive       bool      `json:"is_active"`
	Balance        float64   `json:"balance"`
	Delegated      float64   `json:"delegated"`
	OwnStake       float64   `json:"own_stake"`
	StakingBalance float64   `json:"staking_balance"`
	NDelegations   int64     `json:"n_delegations"`
	NStakers       int64     `json:"n_stakers"`
	Since          int64     `json:"since"`
	SinceTime      time.Time `json:"since_time"`
}

type StakeSnapshotList []*Snapshot

type StakeSnapshotQuery = client.TableQuery[*StakeSnapshot]

func (c bakerClient) NewStakeSnapshotQuery() *StakeSnapshotQuery {
	return client.NewTableQuery[*StakeSnapshot](c.client, "snapshot")
}
