// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type LendingPosition struct {
	Id                 uint64     `json:"id"`
	Contract           string     `json:"contract"`
	PoolId             int64      `json:"pool_id"`
	Name               string     `json:"name"`
	Entity             string     `json:"entity"`
	Owner              string     `json:"owner"`
	StakeId            int64      `json:"stake_id"`
	DebtToken          string     `json:"debt_token"`
	CollateralToken    string     `json:"collateral_token"`
	DebtDecimals       int        `json:"debt_decimals"`
	CollateralDecimals int        `json:"collateral_decimals"`
	DebtSymbol         string     `json:"debt_symbol"`
	CollateralSymbol   string     `json:"collateral_symbol"`
	Balance            tezos.Z    `json:"balance"`
	Deposited          tezos.Z    `json:"deposited"`
	Withdrawn          tezos.Z    `json:"withdrawn"`
	Borrowed           tezos.Z    `json:"borrowed"`
	Repaid             tezos.Z    `json:"repaid"`
	Liquidated         tezos.Z    `json:"liquidated"`
	Sent               tezos.Z    `json:"sent"`
	Received           tezos.Z    `json:"received"`
	InterestEarned     tezos.Z    `json:"interest_earned"`
	InterestPaid       tezos.Z    `json:"interest_paid"`
	InterestPending    tezos.Z    `json:"interest_pending"`
	IsClosed           bool       `json:"is_closed"`
	OpenBlock          int64      `json:"open_block"`
	OpenTime           time.Time  `json:"open_time"`
	CloseBlock         int64      `json:"close_block,omitempty"`
	CloseTime          *time.Time `json:"close_time,omitempty"`
	PositionUSD        float64    `json:"position_value_usd,string"`
	InterestUSD        float64    `json:"interest_pending_usd,string"`
}

type LendingPositionParams = Params[LendingPosition]

func NewLendingPositionParams() LendingPositionParams {
	return LendingPositionParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) ListLendingPositions(ctx context.Context, params LendingPositionParams) ([]*LendingPosition, error) {
	list := make([]*LendingPosition, 0)
	u := params.WithPath("/v1/lend/positions").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListLendingPoolPositions(ctx context.Context, addr PoolAddress, params LendingPositionParams) ([]*LendingPosition, error) {
	list := make([]*LendingPosition, 0)
	u := params.WithPath(fmt.Sprintf("/v1/lend/%s/positions", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListWalletLendingPositions(ctx context.Context, addr tezos.Address, params LendingPositionParams) ([]*LendingPosition, error) {
	list := make([]*LendingPosition, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/lend_positions", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
