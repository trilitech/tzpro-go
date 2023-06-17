// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type LendingEvent struct {
	Id                 uint64    `json:"id"`
	Contract           string    `json:"contract"`
	PoolId             int64     `json:"pool_id"`
	Name               string    `json:"name"`
	Entity             string    `json:"entity"`
	Type               string    `json:"event_type"`
	DebtToken          string    `json:"debt_token"`
	CollateralToken    string    `json:"collateral_token"`
	DebtDecimals       int       `json:"debt_decimals"`
	CollateralDecimals int       `json:"collateral_decimals"`
	DebtSymbol         string    `json:"debt_symbol"`
	CollateralSymbol   string    `json:"collateral_symbol"`
	Owner              string    `json:"owner"`
	StakeId            int64     `json:"stake_id"`
	Volume             tezos.Z   `json:"volume"`
	Debt               tezos.Z   `json:"debt"`
	Collateral         tezos.Z   `json:"collateral"`
	Fee                tezos.Z   `json:"fee"`
	Interest           tezos.Z   `json:"interest"`
	Signer             string    `json:"signer"`
	Sender             string    `json:"sender"`
	Receiver           string    `json:"receiver"`
	TxHash             string    `json:"tx_hash"`
	TxFee              int64     `json:"tx_fee,string"`
	Block              int64     `json:"block"`
	Time               time.Time `json:"time"`
}

type LendingEventParams = Params[LendingEvent]

func NewLendingEventParams() LendingEventParams {
	return LendingEventParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) ListLendingEvents(ctx context.Context, params LendingEventParams) ([]*LendingEvent, error) {
	list := make([]*LendingEvent, 0)
	u := params.WithPath("/v1/lend/events").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListLendingPoolEvents(ctx context.Context, addr PoolAddress, params LendingEventParams) ([]*LendingEvent, error) {
	list := make([]*LendingEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/lend/%s/events", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListWalletLendingEvents(ctx context.Context, addr tezos.Address, params LendingEventParams) ([]*LendingEvent, error) {
	list := make([]*LendingEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/lend_events", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
