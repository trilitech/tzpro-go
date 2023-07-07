// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"
	"fmt"
	"time"
)

type LendingPosition struct {
	Id                 uint64       `json:"id"`
	Contract           Address      `json:"contract"`
	PoolId             int          `json:"pool_id"`
	Name               string       `json:"name"`
	Entity             string       `json:"entity"`
	Owner              Address      `json:"owner"`
	StakeId            int64        `json:"stake_id"`
	DebtToken          TokenAddress `json:"debt_token"`
	CollateralToken    TokenAddress `json:"collateral_token"`
	DebtDecimals       int          `json:"debt_decimals"`
	CollateralDecimals int          `json:"collateral_decimals"`
	DebtSymbol         string       `json:"debt_symbol"`
	CollateralSymbol   string       `json:"collateral_symbol"`
	Balance            Z            `json:"balance"`
	Deposited          Z            `json:"deposited"`
	Withdrawn          Z            `json:"withdrawn"`
	Borrowed           Z            `json:"borrowed"`
	Repaid             Z            `json:"repaid"`
	Liquidated         Z            `json:"liquidated"`
	Sent               Z            `json:"sent"`
	Received           Z            `json:"received"`
	InterestEarned     Z            `json:"interest_earned"`
	InterestPaid       Z            `json:"interest_paid"`
	InterestPending    Z            `json:"interest_pending"`
	IsClosed           bool         `json:"is_closed"`
	OpenBlock          int64        `json:"open_block"`
	OpenTime           time.Time    `json:"open_time"`
	CloseBlock         int64        `json:"close_block,omitempty"`
	CloseTime          *time.Time   `json:"close_time,omitempty"`
	PositionUSD        float64      `json:"position_value_usd,string"`
	InterestUSD        float64      `json:"interest_pending_usd,string"`
}

func (c *lendClient) ListPositions(ctx context.Context, params Query) ([]*LendingPosition, error) {
	list := make([]*LendingPosition, 0)
	u := params.WithPath("/v1/lend/positions").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *lendClient) ListPoolPositions(ctx context.Context, addr PoolAddress, params Query) ([]*LendingPosition, error) {
	list := make([]*LendingPosition, 0)
	u := params.WithPath(fmt.Sprintf("/v1/lend/%s/positions", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
