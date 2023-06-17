// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
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
	Volume         tezos.Z   `json:"volume"`
	Reward         tezos.Z   `json:"reward"`
	Fee            tezos.Z   `json:"fee"`
	FeeBps         float64   `json:"fee_bps,string"`
	StakeId        int64     `json:"stake_id"`
	StakeSupply    tezos.Z   `json:"stake_supply"`
	RewardSupply   tezos.Z   `json:"reward_supply"`
	Signer         string    `json:"signer"`
	Sender         string    `json:"sender"`
	Receiver       string    `json:"receiver"`
	TxHash         string    `json:"tx_hash"`
	TxFee          int64     `json:"tx_fee,string"`
	Block          int64     `json:"block"`
	Time           time.Time `json:"time"`
}

type FarmEventParams = Params[FarmEvent]

func NewFarmEventParams() FarmEventParams {
	return FarmEventParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) ListFarmEvents(ctx context.Context, params FarmEventParams) ([]*FarmEvent, error) {
	list := make([]*FarmEvent, 0)
	u := params.WithPath("/v1/farm/events").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListFarmPoolEvents(ctx context.Context, addr PoolAddress, params FarmEventParams) ([]*FarmEvent, error) {
	list := make([]*FarmEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/farm/%s/events", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListWalletFarmEvents(ctx context.Context, addr tezos.Address, params FarmEventParams) ([]*FarmEvent, error) {
	list := make([]*FarmEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/farm_events", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
