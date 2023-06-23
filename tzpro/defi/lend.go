// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzpro-go/internal/client"
)

type LendingAPI interface {
	GetPool(context.Context, PoolAddress) (*LendingPool, error)
	ListPoolEvents(context.Context, PoolAddress, Params) ([]*LendingEvent, error)
	ListPoolPositions(context.Context, PoolAddress, Params) ([]*LendingPosition, error)

	// firehose
	ListPools(context.Context, Params) ([]*LendingPool, error)
	ListEvents(context.Context, Params) ([]*LendingEvent, error)
	ListPositions(context.Context, Params) ([]*LendingPosition, error)
}

func NewLendingAPI(c *client.Client) LendingAPI {
	return &lendClient{client: c}
}

type lendClient struct {
	client *client.Client
}

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

func (p LendingPool) Address() PoolAddress {
	a, _ := ParseAddress(p.Contract)
	return NewPoolAddress(a, p.PoolId)
}

func (c *lendClient) GetPool(ctx context.Context, addr PoolAddress) (*LendingPool, error) {
	p := &LendingPool{}
	u := fmt.Sprintf("/v1/lend/%s", addr)
	if err := c.client.Get(ctx, u, nil, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (c *lendClient) ListPools(ctx context.Context, params Params) ([]*LendingPool, error) {
	list := make([]*LendingPool, 0)
	u := params.WithPath("/v1/lend").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
