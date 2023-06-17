// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"context"
	// "encoding/json"
	"fmt"
	// "strconv"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type Income struct {
	RowId                  uint64        `json:"row_id"`
	Cycle                  int64         `json:"cycle"`
	Address                tezos.Address `json:"address"`
	AccountId              uint64        `json:"account_id"`
	Rolls                  int64         `json:"rolls"`
	Balance                float64       `json:"balance"`
	Delegated              float64       `json:"delegated"`
	ActiveStake            float64       `json:"active_stake"`
	NDelegations           int64         `json:"n_delegations"`
	NBakingRights          int64         `json:"n_baking_rights"`
	NEndorsingRights       int64         `json:"n_endorsing_rights"`
	Luck                   float64       `json:"luck"`
	LuckPct                float64       `json:"luck_percent"`
	ContributionPct        float64       `json:"contribution_percent"`
	PerformancePct         float64       `json:"performance_percent"`
	NBlocksBaked           int64         `json:"n_blocks_baked"`
	NBlocksProposed        int64         `json:"n_blocks_proposed"`
	NBlocksNotBaked        int64         `json:"n_blocks_not_baked"`
	NBlocksEndorsed        int64         `json:"n_blocks_endorsed"`
	NBlocksNotEndorsed     int64         `json:"n_blocks_not_endorsed"`
	NSlotsEndorsed         int64         `json:"n_slots_endorsed"`
	NSeedsRevealed         int64         `json:"n_seeds_revealed"`
	ExpectedIncome         float64       `json:"expected_income"`
	TotalIncome            float64       `json:"total_income"`
	TotalDeposits          float64       `json:"total_deposits"`
	BakingIncome           float64       `json:"baking_income"`
	EndorsingIncome        float64       `json:"endorsing_income"`
	AccusationIncome       float64       `json:"accusation_income"`
	SeedIncome             float64       `json:"seed_income"`
	FeesIncome             float64       `json:"fees_income"`
	TotalLoss              float64       `json:"total_loss"`
	AccusationLoss         float64       `json:"accusation_loss"`
	SeedLoss               float64       `json:"seed_loss"`
	EndorsingLoss          float64       `json:"endorsing_loss"`
	LostAccusationFees     float64       `json:"lost_accusation_fees"`
	LostAccusationRewards  float64       `json:"lost_accusation_rewards"`
	LostAccusationDeposits float64       `json:"lost_accusation_deposits"`
	LostSeedFees           float64       `json:"lost_seed_fees"`
	LostSeedRewards        float64       `json:"lost_seed_rewards"`
	StartTime              time.Time     `json:"start_time"`
	EndTime                time.Time     `json:"end_time"`
}

type IncomeList struct {
	Rows    []*Income
	columns []string
}

func (l IncomeList) Len() int {
	return len(l.Rows)
}

func (l IncomeList) Cursor() uint64 {
	if len(l.Rows) == 0 {
		return 0
	}
	return l.Rows[len(l.Rows)-1].RowId
}

func (l *IncomeList) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, null) {
		return nil
	}
	if data[0] != '[' {
		return fmt.Errorf("IncomeList: expected JSON array")
	}
	return DecodeSlice(data, l.columns, &l.Rows)
}

type IncomeQuery struct {
	tableQuery
}

func (c *Client) NewIncomeQuery() IncomeQuery {
	return IncomeQuery{c.newTableQuery("income", &Income{})}
}

func (q IncomeQuery) Run(ctx context.Context) (*IncomeList, error) {
	result := &IncomeList{
		columns: q.Columns,
	}
	if err := q.client.QueryTable(ctx, &q.tableQuery, result); err != nil {
		return nil, err
	}
	return result, nil
}
