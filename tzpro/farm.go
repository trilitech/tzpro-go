// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

//nolint:staticcheck
type Farm struct {
	Id               uint64     `json:"id"`
	Contract         string     `json:"contract"`
	PoolId           int        `json:"pool_id"`
	Creator          string     `json:"creator"`
	Name             string     `json:"name"`
	Entity           string     `json:"entity"`
	StakeToken       *Token     `json:"stake_token"`
	RewardToken      *Token     `json:"reward_token"`
	FirstBlock       int64      `json:"first_block"`
	FirstTime        time.Time  `json:"first_time"`
	Tags             []string   `json:"tags"`
	TotalStake       string     `json:"total_stake"`
	RemainingRewards string     `json:"remaining_rewards"`
	NumPositions     int        `json:"num_positions"`
	StartTime        *time.Time `json:"start_time,omitempty"`
	EndTime          *time.Time `json:"end_time,omitempty"`
}

type FarmParams = Params[Farm]

func NewFarmParams() FarmParams {
	return FarmParams{
		Query: make(map[string][]string),
	}
}

func (p Farm) Address() PoolAddress {
	a, _ := tezos.ParseAddress(p.Contract)
	return NewPoolAddress(a, p.PoolId)
}

func (c *Client) GetFarm(ctx context.Context, addr PoolAddress, params FarmParams) (*Farm, error) {
	p := &Farm{}
	u := params.WithPath(fmt.Sprintf("/v1/farm/%s", addr)).Url()
	if err := c.get(ctx, u, nil, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (c *Client) ListFarms(ctx context.Context, params FarmParams) ([]*Farm, error) {
	list := make([]*Farm, 0)
	u := params.WithPath("/v1/farm").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
