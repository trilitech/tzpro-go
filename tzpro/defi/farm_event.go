// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"
	"fmt"
	"time"
)

type FarmEvent struct {
	Id             uint64    `json:"id"`
	Contract       string    `json:"contract"`
	PoolId         int64     `json:"pool_id"`
	Name           string    `json:"name"`
	Entity         string    `json:"entity"`
	Type           string    `json:"event_type"`
	StakeToken     string    `json:"stake_token"`
	RewardToken    string    `json:"reward_token"`
	StakeSymbol    string    `json:"stake_symbol"`
	RewardSymbol   string    `json:"reward_symbol"`
	StakeDecimals  int       `json:"stake_decimals"`
	RewardDecimals int       `json:"reward_decimals"`
	Volume         Z         `json:"volume"`
	Reward         Z         `json:"reward"`
	Fee            Z         `json:"fee"`
	FeeBps         float64   `json:"fee_bps,string"`
	StakeId        int64     `json:"stake_id"`
	StakeSupply    Z         `json:"stake_supply"`
	RewardSupply   Z         `json:"reward_supply"`
	Signer         string    `json:"signer"`
	Sender         string    `json:"sender"`
	Receiver       string    `json:"receiver"`
	TxHash         string    `json:"tx_hash"`
	TxFee          int64     `json:"tx_fee,string"`
	Block          int64     `json:"block"`
	Time           time.Time `json:"time"`
}

func (c *farmClient) ListEvents(ctx context.Context, params Query) ([]*FarmEvent, error) {
	list := make([]*FarmEvent, 0)
	u := params.WithPath("/v1/farm/events").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *farmClient) ListPoolEvents(ctx context.Context, addr PoolAddress, params Query) ([]*FarmEvent, error) {
	list := make([]*FarmEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/farm/%s/events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
