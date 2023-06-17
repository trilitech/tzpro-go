// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type LendingPool struct {
	Id              uint64    `json:"id"`
	Contract        string    `json:"contract"`
	PoolId          int       `json:"pool_id"`
	Creator         string    `json:"creator"`
	Name            string    `json:"name"`
	Entity          string    `json:"entity"`
	DebtToken       *Token    `json:"debt_token"`
	CollateralToken *Token    `json:"collateral_token"`
	FirstBlock      int64     `json:"first_block"`
	FirstTime       time.Time `json:"first_time"`
	Tags            []string  `json:"tags"`
	TotalCollateral string    `json:"total_collateral"`
	TotalDebt       string    `json:"total_debt"`
	NumDeposits     int       `json:"num_deposits"`
	NumBorrows      int       `json:"num_borrows"`
}

type LendingPoolParams = Params[LendingPool]

func NewLendingPoolParams() LendingPoolParams {
	return LendingPoolParams{
		Query: make(map[string][]string),
	}
}

func (p LendingPool) Address() PoolAddress {
	a, _ := tezos.ParseAddress(p.Contract)
	return NewPoolAddress(a, p.PoolId)
}

func (c *Client) GetLendingPool(ctx context.Context, addr PoolAddress, params LendingPoolParams) (*LendingPool, error) {
	p := &LendingPool{}
	u := params.WithPath(fmt.Sprintf("/v1/lend/%s", addr)).Url()
	if err := c.get(ctx, u, nil, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (c *Client) ListLendingPools(ctx context.Context, params LendingPoolParams) ([]*LendingPool, error) {
	list := make([]*LendingPool, 0)
	u := params.WithPath("/v1/lend").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
