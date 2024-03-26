// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"
	"fmt"
	"time"
)

type Election struct {
	Id                int       `json:"election_id"`
	MaxPeriods        int       `json:"max_periods"`
	NumPeriods        int       `json:"num_periods"`
	NumProposals      int       `json:"num_proposals"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	StartHeight       int64     `json:"start_height"`
	EndHeight         int64     `json:"end_height"`
	IsEmpty           bool      `json:"is_empty"`
	IsOpen            bool      `json:"is_open"`
	IsFailed          bool      `json:"is_failed"`
	NoQuorum          bool      `json:"no_quorum"`
	NoMajority        bool      `json:"no_majority"`
	NoProposal        bool      `json:"no_proposal"`
	VotingPeriodKind  string    `json:"voting_period"`
	ProposalPeriod    *Vote     `json:"proposal"`
	ExplorationPeriod *Vote     `json:"exploration"`
	CooldownPeriod    *Vote     `json:"cooldown"`
	PromotionPeriod   *Vote     `json:"promotion"`
	AdoptionPeriod    *Vote     `json:"adoption"`
}

func (e Election) Period(p string) *Vote {
	switch p {
	case "proposal":
		return e.ProposalPeriod
	case "exploration":
		return e.ExplorationPeriod
	case "cooldown":
		return e.CooldownPeriod
	case "promotion":
		return e.PromotionPeriod
	case "adoption":
		return e.AdoptionPeriod
	default:
		return nil
	}
}

type Vote struct {
	VotingPeriod     int64       `json:"voting_period"`
	VotingPeriodKind string      `json:"voting_period_kind"`
	StartTime        time.Time   `json:"period_start_time"`
	EndTime          time.Time   `json:"period_end_time"`
	StartHeight      int64       `json:"period_start_block"`
	EndHeight        int64       `json:"period_end_block"`
	EligibleStake    float64     `json:"eligible_stake"`
	EligibleVoters   int         `json:"eligible_voters"`
	QuorumPct        int         `json:"quorum_pct"`
	QuorumStake      float64     `json:"quorum_stake"`
	TurnoutStake     float64     `json:"turnout_stake"`
	TurnoutVoters    int         `json:"turnout_voters"`
	TurnoutPct       int         `json:"turnout_pct"`
	TurnoutEma       int         `json:"turnout_ema"`
	YayStake         float64     `json:"yay_stake"`
	YayVoters        int         `json:"yay_voters"`
	NayStake         float64     `json:"nay_stake"`
	NayVoters        int         `json:"nay_voters"`
	PassStake        float64     `json:"pass_stake"`
	PassVoters       int         `json:"pass_voters"`
	IsOpen           bool        `json:"is_open"`
	IsFailed         bool        `json:"is_failed"`
	IsDraw           bool        `json:"is_draw"`
	NoProposal       bool        `json:"no_proposal"`
	NoQuorum         bool        `json:"no_quorum"`
	NoMajority       bool        `json:"no_majority"`
	Proposals        []*Proposal `json:"proposals"`
}

type Proposal struct {
	Hash          string    `json:"hash"`
	SourceAddress Address   `json:"source"`
	OpId          uint64    `json:"op"`
	Height        int64     `json:"height"`
	Time          time.Time `json:"time"`
	Stake         float64   `json:"stake"`
	Voters        int64     `json:"voters"`
}

type Ballot struct {
	Id               uint64    `json:"id"`
	Height           int64     `json:"height"`
	Timestamp        time.Time `json:"time"`
	OpHash           uint64    `json:"op"`
	ElectionId       int       `json:"election_id"`
	VotingPeriod     int64     `json:"voting_period"`
	VotingPeriodKind string    `json:"voting_period_kind"`
	Proposal         string    `json:"proposal"`
	Ballot           string    `json:"ballot"`
	Stake            float64   `json:"stake"`
	Sender           Address   `json:"sender"`
}

type BallotList []*Ballot

type Voter struct {
	Id        uint64   `json:"id"`
	Address   Address  `json:"address"`
	Stake     float64  `json:"stake"`
	Ballot    string   `json:"ballot"`
	HasVoted  bool     `json:"has_voted"`
	Proposals []string `json:"proposals"`
}

func (c *explorerClient) GetElection(ctx context.Context, id int) (*Election, error) {
	e := &Election{}
	u := fmt.Sprintf("/explorer/election/%d", id)
	if err := c.client.Get(ctx, u, nil, e); err != nil {
		return nil, err
	}
	return e, nil
}

func (c *explorerClient) ListVoters(ctx context.Context, id int, stage int) ([]Voter, error) {
	voters := make([]Voter, 0)
	u := fmt.Sprintf("/explorer/election/%d/%d/voters?limit=5000", id, stage)
	if err := c.client.Get(ctx, u, nil, &voters); err != nil {
		return nil, err
	}
	return voters, nil
}

func (c *explorerClient) ListBallots(ctx context.Context, id int, stage int) (BallotList, error) {
	ballots := make(BallotList, 0)
	u := fmt.Sprintf("/explorer/election/%d/%d/ballots?limit=5000", id, stage)
	if err := c.client.Get(ctx, u, nil, &ballots); err != nil {
		return nil, err
	}
	return ballots, nil
}
