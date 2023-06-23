// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"
	"time"
)

type Tip struct {
	Name               string    `json:"name"`
	Network            string    `json:"network"`
	Symbol             string    `json:"symbol"`
	ChainId            string    `json:"chain_id"`
	GenesisTime        time.Time `json:"genesis_time"`
	Hash               BlockHash `json:"block_hash"`
	Height             int64     `json:"height"`
	Cycle              int64     `json:"cycle"`
	Timestamp          time.Time `json:"timestamp"`
	Protocol           string    `json:"protocol"`
	TotalAccounts      int64     `json:"total_accounts"`
	TotalContracts     int64     `json:"total_contracts"`
	TotalRollups       int64     `json:"total_rollups"`
	FundedAccounts     int64     `json:"funded_accounts"`
	DustAccounts       int64     `json:"dust_accounts"`
	DustDelegators     int64     `json:"dust_delegators"`
	TotalOps           int64     `json:"total_ops"`
	Delegators         int64     `json:"delegators"`
	Bakers             int64     `json:"bakers"`
	Rolls              int64     `json:"rolls"`
	RollOwners         int64     `json:"roll_owners"`
	NewAccounts30d     int64     `json:"new_accounts_30d"`
	ClearedAccounts30d int64     `json:"cleared_accounts_30d"`
	FundedAccounts30d  int64     `json:"funded_accounts_30d"`
	Inflation1Y        float64   `json:"inflation_1y"`
	InflationRate1Y    float64   `json:"inflation_rate_1y"`
	Health             int       `json:"health"`
	Supply             *Supply   `json:"supply,omitempty"`
	Status             Status    `json:"status"`
}

func (c *explorerClient) GetTip(ctx context.Context) (*Tip, error) {
	tip := &Tip{}
	if err := c.client.Get(ctx, "/explorer/tip", nil, tip); err != nil {
		return nil, err
	}
	return tip, nil
}

type Supply struct {
	RowId               uint64    `json:"row_id"`
	Height              int64     `json:"height"`
	Cycle               int64     `json:"cycle"`
	Timestamp           time.Time `json:"time"`
	Total               float64   `json:"total"`
	Activated           float64   `json:"activated"`
	Unclaimed           float64   `json:"unclaimed"`
	Circulating         float64   `json:"circulating"`
	Liquid              float64   `json:"liquid"`
	Delegated           float64   `json:"delegated"`
	Staking             float64   `json:"staking"`
	Shielded            float64   `json:"shielded"`
	ActiveStake         float64   `json:"active_stake"`
	ActiveDelegated     float64   `json:"active_delegated"`
	ActiveStaking       float64   `json:"active_staking"`
	InactiveDelegated   float64   `json:"inactive_delegated"`
	InactiveStaking     float64   `json:"inactive_staking"`
	Minted              float64   `json:"minted"`
	MintedBaking        float64   `json:"minted_baking"`
	MintedEndorsing     float64   `json:"minted_endorsing"`
	MintedSeeding       float64   `json:"minted_seeding"`
	MintedAirdrop       float64   `json:"minted_airdrop"`
	MintedSubsidy       float64   `json:"minted_subsidy"`
	Burned              float64   `json:"burned"`
	BurnedDoubleBaking  float64   `json:"burned_double_baking"`
	BurnedDoubleEndorse float64   `json:"burned_double_endorse"`
	BurnedOrigination   float64   `json:"burned_origination"`
	BurnedAllocation    float64   `json:"burned_allocation"`
	BurnedStorage       float64   `json:"burned_storage"`
	BurnedExplicit      float64   `json:"burned_explicit"`
	BurnedSeedMiss      float64   `json:"burned_seed_miss"`
	BurnedAbsence       float64   `json:"burned_absence"`
	BurnedRollup        float64   `json:"burned_rollup"`
	Frozen              float64   `json:"frozen"`
	FrozenDeposits      float64   `json:"frozen_deposits"`
	FrozenRewards       float64   `json:"frozen_rewards"`
	FrozenFees          float64   `json:"frozen_fees"`
	FrozenBonds         float64   `json:"frozen_bonds"`
}
