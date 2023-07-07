// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzpro-go/internal/client"
)

type FarmAPI interface {
	GetFarm(context.Context, PoolAddress) (*Farm, error)
	ListPoolEvents(context.Context, PoolAddress, Query) ([]*FarmEvent, error)
	ListFarmPoolPositions(context.Context, PoolAddress, Query) ([]*FarmPosition, error)

	// firehose
	ListFarms(context.Context, Query) ([]*Farm, error)
	ListEvents(context.Context, Query) ([]*FarmEvent, error)
	ListPositions(context.Context, Query) ([]*FarmPosition, error)
}

func NewFarmAPI(c *client.Client) FarmAPI {
	return &farmClient{client: c}
}

type farmClient struct {
	client *client.Client
}

type Farm struct {
	Id               uint64     `json:"id"`
	Contract         Address    `json:"contract"`
	PoolId           int        `json:"pool_id"`
	Creator          Address    `json:"creator"`
	Name             string     `json:"name"`
	Entity           string     `json:"entity"`
	StakeToken       *Token     `json:"stake_token"`
	RewardToken      *Token     `json:"reward_token"`
	FirstBlock       int64      `json:"first_block"`
	FirstTime        time.Time  `json:"first_time"`
	Tags             []string   `json:"tags"`
	TotalStake       Z          `json:"total_stake"`
	RemainingRewards Z          `json:"remaining_rewards"`
	NumPositions     int        `json:"num_positions"`
	StartTime        *time.Time `json:"start_time,omitempty"`
	EndTime          *time.Time `json:"end_time,omitempty"`
}

func (p Farm) Address() PoolAddress {
	return NewPoolAddress(p.Contract, p.PoolId)
}

func (c *farmClient) GetFarm(ctx context.Context, addr PoolAddress) (*Farm, error) {
	p := &Farm{}
	u := fmt.Sprintf("/v1/farm/%s", addr)
	if err := c.client.Get(ctx, u, nil, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (c *farmClient) ListFarms(ctx context.Context, params Query) ([]*Farm, error) {
	list := make([]*Farm, 0)
	u := params.WithPath("/v1/farm").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
