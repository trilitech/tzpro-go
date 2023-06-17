// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"context"
	"fmt"
	"time"
)

type Chain struct {
	RowId                uint64    `json:"row_id"`
	Height               int64     `json:"height"`
	Cycle                int64     `json:"cycle"`
	Timestamp            time.Time `json:"time"`
	TotalAccounts        int64     `json:"total_accounts"`
	TotalContracts       int64     `json:"total_contracts"`
	TotalRollups         int64     `json:"total_rollups"`
	TotalOps             int64     `json:"total_ops"`
	TotalOpsFailed       int64     `json:"total_ops_failed"`
	TotalContractOps     int64     `json:"total_contract_ops"`
	TotalContractCalls   int64     `json:"total_contract_calls"`
	TotalRollupCalls     int64     `json:"total_rollup_calls"`
	TotalActivations     int64     `json:"total_activations"`
	TotalNonces          int64     `json:"total_nonce_revelations"`
	TotalEndorsements    int64     `json:"total_endorsements"`
	TotalPreendorsements int64     `json:"total_preendorsements"`
	TotalDoubleBake      int64     `json:"total_double_bakings"`
	TotalDoubleEndorse   int64     `json:"total_double_endorsements"`
	TotalDelegations     int64     `json:"total_delegations"`
	TotalReveals         int64     `json:"total_reveals"`
	TotalOriginations    int64     `json:"total_originations"`
	TotalTransactions    int64     `json:"total_transactions"`
	TotalProposals       int64     `json:"total_proposals"`
	TotalBallots         int64     `json:"total_ballots"`
	TotalConstants       int64     `json:"total_constants"`
	TotalSetLimits       int64     `json:"total_set_limits"`
	TotalStorageBytes    int64     `json:"total_storage_bytes"`
	TotalTicketTransfers int64     `json:"total_ticket_transfers"`
	FundedAccounts       int64     `json:"funded_accounts"`
	DustAccounts         int64     `json:"dust_accounts"`
	GhostAccounts        int64     `json:"ghost_accounts"`
	UnclaimedAccounts    int64     `json:"unclaimed_accounts"`
	TotalDelegators      int64     `json:"total_delegators"`
	ActiveDelegators     int64     `json:"active_delegators"`
	InactiveDelegators   int64     `json:"inactive_delegators"`
	DustDelegators       int64     `json:"dust_delegators"`
	TotalBakers          int64     `json:"total_bakers"`
	ActiveBakers         int64     `json:"active_bakers"`
	InactiveBakers       int64     `json:"inactive_bakers"`
	ZeroBakers           int64     `json:"zero_bakers"`
	SelfBakers           int64     `json:"self_bakers"`
	SingleBakers         int64     `json:"single_bakers"`
	MultiBakers          int64     `json:"multi_bakers"`
	Rolls                int64     `json:"rolls"`
	RollOwners           int64     `json:"roll_owners"`
}

type ChainList struct {
	Rows    []*Chain
	columns []string
}

func (l ChainList) Len() int {
	return len(l.Rows)
}

func (l ChainList) Cursor() uint64 {
	if len(l.Rows) == 0 {
		return 0
	}
	return l.Rows[len(l.Rows)-1].RowId
}

func (l *ChainList) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, null) {
		return nil
	}
	if data[0] != '[' {
		return fmt.Errorf("ChainList: expected JSON array")
	}
	return DecodeSlice(data, l.columns, &l.Rows)
}

type ChainQuery struct {
	tableQuery
}

func (c *Client) NewChainQuery() ChainQuery {
	return ChainQuery{c.newTableQuery("chain", &Chain{})}
}

func (q ChainQuery) Run(ctx context.Context) (*ChainList, error) {
	result := &ChainList{
		columns: q.Columns,
	}
	if err := q.client.QueryTable(ctx, &q.tableQuery, result); err != nil {
		return nil, err
	}
	return result, nil
}
