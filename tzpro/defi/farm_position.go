// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"
	"fmt"
	"time"
)

type FarmPosition struct {
	Id             uint64     `json:"id"`
	Contract       string     `json:"contract"`
	PoolId         int64      `json:"pool_id"`
	Name           string     `json:"name"`
	Entity         string     `json:"entity"`
	Owner          string     `json:"owner"`
	StakeToken     string     `json:"stake_token"`
	RewardToken    string     `json:"reward_token"`
	StakeSymbol    string     `json:"stake_symbol"`
	RewardSymbol   string     `json:"reward_symbol"`
	StakeDecimals  int        `json:"stake_decimals"`
	RewardDecimals int        `json:"reward_decimals"`
	TotalStake     string     `json:"total_stake"`
	StakeId        int64      `json:"stake_id"`
	Stake          Z          `json:"stake_balance"`
	Deposited      Z          `json:"stake_deposited"`
	Withdrawn      Z          `json:"stake_withdrawn"`
	Claimed        Z          `json:"rewards_claimed"`
	Pending        Z          `json:"rewards_pending"`
	IsClosed       bool       `json:"is_closed"`
	OpenBlock      int64      `json:"open_block"`
	OpenTime       time.Time  `json:"open_time"`
	CloseBlock     int64      `json:"close_block,omitempty"`
	CloseTime      *time.Time `json:"close_time,omitempty"`
}

func (c *farmClient) ListPositions(ctx context.Context, params Query) ([]*FarmPosition, error) {
	list := make([]*FarmPosition, 0)
	u := params.WithPath("/v1/farm/positions").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *farmClient) ListFarmPoolPositions(ctx context.Context, addr PoolAddress, params Query) ([]*FarmPosition, error) {
	list := make([]*FarmPosition, 0)
	u := params.WithPath(fmt.Sprintf("/v1/farm/%s/positions", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
