// Copyright (c) 2020-2024 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"
	"fmt"
	"time"
)

type Tip struct {
	Name        string       `json:"name"`
	Network     string       `json:"network"`
	Symbol      string       `json:"symbol"`
	ChainId     ChainIdHash  `json:"chain_id"`
	GenesisTime time.Time    `json:"genesis_time"`
	Block       BlockHash    `json:"block"`
	Height      int64        `json:"height"`
	Cycle       int64        `json:"cycle"`
	Time        time.Time    `json:"time"`
	Protocol    ProtocolHash `json:"protocol"`

	Stats  *Statistics `json:"stats,omitempty"`
	Supply *Supply     `json:"supply,omitempty"`
	Totals *Chain      `json:"totals,omitempty"`
	Status *Status     `json:"status,omitempty"`
}

func (c *explorerClient) GetTip(ctx context.Context) (*Tip, error) {
	tip := &Tip{}
	if err := c.client.Get(ctx, "/explorer/tip", nil, tip); err != nil {
		return nil, err
	}
	return tip, nil
}

type Statistics struct {
	Time                 time.Time `json:"time"`
	NewAccounts30d       int64     `json:"new_accounts_30d"`
	NewContracts30d      int64     `json:"new_contracts_30d"`
	NewFundedAccounts30d int64     `json:"new_funded_accounts_30d"`
	NewGhostAccounts30d  int64     `json:"new_ghost_accounts_30d"`
	ContractCalls30d     int64     `json:"contract_calls_30d"`
	Inflation1Y          float64   `json:"inflation_1y"`
	InflationRate1Y      float64   `json:"inflation_rate_1y"`
}

type Supply struct {
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
	Unstaking           float64   `json:"unstaking"`
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
	BurnedAbsence       float64   `json:"burned_offline"`
	BurnedRollup        float64   `json:"burned_rollup"`
	Frozen              float64   `json:"frozen"`
	FrozenDeposits      float64   `json:"frozen_deposits"`
	FrozenRewards       float64   `json:"frozen_rewards"`
	FrozenFees          float64   `json:"frozen_fees"`
	FrozenBonds         float64   `json:"frozen_bonds"`
	FrozenStake         float64   `json:"frozen_stake"`
	FrozenBakerStake    float64   `json:"frozen_baker_stake"`
	FrozenStakerStake   float64   `json:"frozen_staker_stake"`
}

func (c *explorerClient) GetSupplyHeight(ctx context.Context, height int64) (*Supply, error) {
	res := &Supply{}
	if err := c.client.Get(ctx, fmt.Sprintf("/explorer/supply/%d", height), nil, res); err != nil {
		return nil, err
	}
	return res, nil
}
