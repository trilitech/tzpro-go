// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzpro-go/internal/client"
)

type BakerAPI interface {
	Get(context.Context, Address, Params) (*Baker, error)
	List(context.Context, Params) (BakerList, error)
	ListVotes(context.Context, Address, Params) (BallotList, error)
	ListEndorsements(context.Context, Address, Params) (OpList, error)
	ListDelegations(context.Context, Address, Params) (OpList, error)
	GetRights(context.Context, Address, int64, Params) (*Rights, error)
	GetIncome(context.Context, Address, int64, Params) (*Income, error)
	GetSnapshot(context.Context, Address, int64, Params) (*Snapshot, error)

	NewIncomeQuery() *IncomeQuery
	NewRightsQuery() *RightsQuery
	NewStakeSnapshotQuery() *StakeSnapshotQuery
}

func NewBakerAPI(c *client.Client) BakerAPI {
	return &bakerClient{client: c}
}

type bakerClient struct {
	client *client.Client
}

type Baker struct {
	Address           Address          `json:"address"`
	BakerSince        time.Time        `json:"baker_since_time"`
	BakerUntil        *time.Time       `json:"baker_until,omitempty"`
	GracePeriod       int64            `json:"grace_period"`
	BakerVersion      string           `json:"baker_version"`
	TotalBalance      float64          `json:"total_balance"`
	SpendableBalance  float64          `json:"spendable_balance"`
	FrozenBalance     float64          `json:"frozen_balance"`
	DelegatedBalance  float64          `json:"delegated_balance"`
	StakingBalance    float64          `json:"staking_balance"`
	StakingCapacity   float64          `json:"staking_capacity"`
	DepositsLimit     *float64         `json:"deposits_limit"`
	StakingShare      float64          `json:"staking_share"`
	ActiveDelegations int64            `json:"active_delegations"`
	IsFull            bool             `json:"is_full"`
	IsActive          bool             `json:"is_active"`
	Events            *BakerEvents     `json:"events,omitempty"`
	Stats             *BakerStatistics `json:"stats,omitempty"`
	Metadata          *Metadata        `json:"metadata,omitempty"`
}

type BakerList []*Baker

type BakerStatistics struct {
	TotalRewardsEarned float64 `json:"total_rewards_earned"`
	TotalFeesEarned    float64 `json:"total_fees_earned"`
	TotalLost          float64 `json:"total_lost"`
	BlocksBaked        int64   `json:"blocks_baked"`
	BlocksProposed     int64   `json:"blocks_proposed"`
	SlotsEndorsed      int64   `json:"slots_endorsed"`
	AvgLuck64          int64   `json:"avg_luck_64"`
	AvgPerformance64   int64   `json:"avg_performance_64"`
	AvgContribution64  int64   `json:"avg_contribution_64"`
	NBakerOps          int64   `json:"n_baker_ops"`
	NProposal          int64   `json:"n_proposals"`
	NBallot            int64   `json:"n_ballots"`
	NEndorsement       int64   `json:"n_endorsements"`
	NPreendorsement    int64   `json:"n_preendorsements"`
	NSeedNonce         int64   `json:"n_nonce_revelations"`
	N2Baking           int64   `json:"n_double_bakings"`
	N2Endorsement      int64   `json:"n_double_endorsements"`
	NSetDepositsLimit  int64   `json:"n_set_limits"`
}

type BakerEvents struct {
	LastBakeHeight    int64     `json:"last_bake_height"`
	LastBakeBlock     string    `json:"last_bake_block"`
	LastBakeTime      time.Time `json:"last_bake_time"`
	LastEndorseHeight int64     `json:"last_endorse_height"`
	LastEndorseBlock  string    `json:"last_endorse_block"`
	LastEndorseTime   time.Time `json:"last_endorse_time"`
	NextBakeHeight    int64     `json:"next_bake_height"`
	NextBakeTime      time.Time `json:"next_bake_time"`
	NextEndorseHeight int64     `json:"next_endorse_height"`
	NextEndorseTime   time.Time `json:"next_endorse_time"`
}

type Delegator struct {
	Address  Address `json:"address"`
	Balance  int64   `json:"balance"`
	IsFunded bool    `json:"is_funded"`
}

type Snapshot struct {
	BakeCycle              int64       `json:"baking_cycle"`
	Height                 int64       `json:"snapshot_height"`
	Cycle                  int64       `json:"snapshot_cycle"`
	Timestamp              time.Time   `json:"snapshot_time"`
	Index                  int         `json:"snapshot_index"`
	Rolls                  int64       `json:"snapshot_rolls"`
	ActiveStake            int64       `json:"active_stake"`
	StakingBalance         int64       `json:"staking_balance"`
	OwnBalance             int64       `json:"own_balance"`
	DelegatedBalance       int64       `json:"delegated_balance"`
	NDelegations           int64       `json:"n_delegations"`
	ExpectedIncome         int64       `json:"expected_income"`
	TotalIncome            int64       `json:"total_income"`
	TotalDeposits          int64       `json:"total_deposits"`
	BakingIncome           int64       `json:"baking_income"`
	EndorsingIncome        int64       `json:"endorsing_income"`
	AccusationIncome       int64       `json:"accusation_income"`
	SeedIncome             int64       `json:"seed_income"`
	FeesIncome             int64       `json:"fees_income"`
	TotalLoss              int64       `json:"total_loss"`
	AccusationLoss         int64       `json:"accusation_loss"`
	SeedLoss               int64       `json:"seed_loss"`
	EndorsingLoss          int64       `json:"endorsing_loss"`
	LostAccusationFees     int64       `json:"lost_accusation_fees"`
	LostAccusationRewards  int64       `json:"lost_accusation_rewards"`
	LostAccusationDeposits int64       `json:"lost_accusation_deposits"`
	LostSeedFees           int64       `json:"lost_seed_fees"`
	LostSeedRewards        int64       `json:"lost_seed_rewards"`
	Delegators             []Delegator `json:"delegators"`
}

func (c *bakerClient) Get(ctx context.Context, addr Address, params Params) (*Baker, error) {
	b := &Baker{}
	u := params.WithPath(fmt.Sprintf("/explorer/bakers/%s", addr)).Url()
	if err := c.client.Get(ctx, u, nil, b); err != nil {
		return nil, err
	}
	return b, nil
}

func (c *bakerClient) List(ctx context.Context, params Params) (BakerList, error) {
	b := make([]*Baker, 0)
	u := params.WithPath("/explorer/bakers").Url()
	if err := c.client.Get(ctx, u, nil, &b); err != nil {
		return nil, err
	}
	return b, nil
}

func (c *bakerClient) ListVotes(ctx context.Context, addr Address, params Params) (BallotList, error) {
	cc := make([]*Ballot, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/bakers/%s/votes", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &cc); err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *bakerClient) ListEndorsements(ctx context.Context, addr Address, params Params) (OpList, error) {
	ops := make(OpList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/bakers/%s/endorsements", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &ops); err != nil {
		return nil, err
	}
	return ops, nil
}

func (c *bakerClient) ListDelegations(ctx context.Context, addr Address, params Params) (OpList, error) {
	ops := make(OpList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/bakers/%s/delegations", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &ops); err != nil {
		return nil, err
	}
	return ops, nil
}

func (c *bakerClient) GetRights(ctx context.Context, addr Address, cycle int64, params Params) (*Rights, error) {
	var r Rights
	u := params.WithPath(fmt.Sprintf("/explorer/bakers/%s/rights/%d", addr, cycle)).Url()
	if err := c.client.Get(ctx, u, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *bakerClient) GetIncome(ctx context.Context, addr Address, cycle int64, params Params) (*Income, error) {
	var r Income
	u := params.WithPath(fmt.Sprintf("/explorer/bakers/%s/income/%d", addr, cycle)).Url()
	if err := c.client.Get(ctx, u, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *bakerClient) GetSnapshot(ctx context.Context, addr Address, cycle int64, params Params) (*Snapshot, error) {
	var r Snapshot
	u := params.WithPath(fmt.Sprintf("/explorer/bakers/%s/snapshot/%d", addr, cycle)).Url()
	if err := c.client.Get(ctx, u, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
