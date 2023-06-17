// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
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
	Stake          tezos.Z    `json:"stake_balance"`
	Deposited      tezos.Z    `json:"stake_deposited"`
	Withdrawn      tezos.Z    `json:"stake_withdrawn"`
	Claimed        tezos.Z    `json:"rewards_claimed"`
	Pending        tezos.Z    `json:"rewards_pending"`
	IsClosed       bool       `json:"is_closed"`
	OpenBlock      int64      `json:"open_block"`
	OpenTime       time.Time  `json:"open_time"`
	CloseBlock     int64      `json:"close_block,omitempty"`
	CloseTime      *time.Time `json:"close_time,omitempty"`
}

type FarmPositionParams = Params[FarmPosition]

func NewFarmPositionParams() FarmPositionParams {
	return FarmPositionParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) ListFarmPositions(ctx context.Context, params FarmPositionParams) ([]*FarmPosition, error) {
	list := make([]*FarmPosition, 0)
	u := params.WithPath("/v1/farm/positions").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListFarmPoolPositions(ctx context.Context, addr PoolAddress, params FarmPositionParams) ([]*FarmPosition, error) {
	list := make([]*FarmPosition, 0)
	u := params.WithPath(fmt.Sprintf("/v1/farm/%s/positions", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListWalletFarmPositions(ctx context.Context, addr tezos.Address, params FarmPositionParams) ([]*FarmPosition, error) {
	list := make([]*FarmPosition, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/farm_positions", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
