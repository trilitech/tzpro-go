// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzpro-go/internal/client"
)

type Chain struct {
	Id                   uint64    `json:"id"`
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
	RollOwners           int64     `json:"eligible_bakers"`
	ActiveBakers         int64     `json:"active_bakers"`
	InactiveBakers       int64     `json:"inactive_bakers"`
	ZeroBakers           int64     `json:"zero_bakers"`
	SelfBakers           int64     `json:"self_bakers"`
	SingleBakers         int64     `json:"single_bakers"`
	MultiBakers          int64     `json:"multi_bakers"`
	TotalStakers         int64     `json:"total_stakers"`
	ActiveStakers        int64     `json:"active_stakers"`
	InactiveStakers      int64     `json:"inactive_stakers"`
}

type ChainQuery = client.TableQuery[*Chain]

func (c *explorerClient) NewChainQuery() *ChainQuery {
	return client.NewTableQuery[*Chain](c.client, "chain")
}

func (c *explorerClient) GetTotalsHeight(ctx context.Context, height int64) (*Chain, error) {
	res := &Chain{}
	if err := c.client.Get(ctx, fmt.Sprintf("/explorer/chain/%d", height), nil, res); err != nil {
		return nil, err
	}
	return res, nil
}
