// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"time"

	"blockwatch.cc/tzpro-go/internal/client"
)

type Income struct {
	RowId                  uint64    `json:"row_id"`     // table only
	Address                Address   `json:"address"`    // table only
	AccountId              uint64    `json:"account_id"` // table only
	Cycle                  int64     `json:"cycle"`
	Rolls                  int64     `json:"snapshot_rolls"    tzpro:"rolls"`
	Balance                float64   `json:"own_balance"       tzpro:"balance"`
	Delegated              float64   `json:"delegated_balance" tzpro:"delegated"`
	Staking                float64   `json:"staking_balance"   tzpro:"-"`
	ActiveStake            float64   `json:"active_stake"`
	NDelegations           int64     `json:"n_delegations"`
	NBakingRights          int64     `json:"n_baking_rights"`
	NEndorsingRights       int64     `json:"n_endorsing_rights"`
	Luck                   float64   `json:"luck"`
	LuckPct                int64     `json:"luck_percent"`
	ContributionPct        int64     `json:"contribution_percent"`
	PerformancePct         int64     `json:"performance_percent"`
	NBlocksBaked           int64     `json:"n_blocks_baked"`
	NBlocksProposed        int64     `json:"n_blocks_proposed"`
	NBlocksNotBaked        int64     `json:"n_blocks_not_baked"`
	NBlocksEndorsed        int64     `json:"n_blocks_endorsed"`
	NBlocksNotEndorsed     int64     `json:"n_blocks_not_endorsed"`
	NSlotsEndorsed         int64     `json:"n_slots_endorsed"`
	NSeedsRevealed         int64     `json:"n_seeds_revealed"`
	ExpectedIncome         float64   `json:"expected_income"`
	TotalIncome            float64   `json:"total_income"`
	TotalDeposits          float64   `json:"total_deposits"`
	BakingIncome           float64   `json:"baking_income"`
	EndorsingIncome        float64   `json:"endorsing_income"`
	AccusationIncome       float64   `json:"accusation_income"`
	SeedIncome             float64   `json:"seed_income"`
	FeesIncome             float64   `json:"fees_income"`
	TotalLoss              float64   `json:"total_loss"`
	AccusationLoss         float64   `json:"accusation_loss"`
	SeedLoss               float64   `json:"seed_loss"`
	EndorsingLoss          float64   `json:"endorsing_loss"`
	LostAccusationFees     float64   `json:"lost_accusation_fees"`
	LostAccusationRewards  float64   `json:"lost_accusation_rewards"`
	LostAccusationDeposits float64   `json:"lost_accusation_deposits"`
	LostSeedFees           float64   `json:"lost_seed_fees"`
	LostSeedRewards        float64   `json:"lost_seed_rewards"`
	StartTime              time.Time `json:"start_time"` // table only
	EndTime                time.Time `json:"end_time"`   // table only
}

type IncomeQuery = client.TableQuery[*Income]

func (c *bakerClient) NewIncomeQuery() *IncomeQuery {
	return client.NewTableQuery[*Income](c.client, "income")
}
