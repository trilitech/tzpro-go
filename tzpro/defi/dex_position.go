// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"
	"fmt"
	"time"
)

type DexPosition struct {
	Id               uint64     `json:"id"`
	Contract         Address    `json:"contract"`
	PairId           int        `json:"pair_id"`
	Name             string     `json:"name"`
	Entity           string     `json:"entity"`
	Pair             string     `json:"pair"`
	Owner            Address    `json:"owner"`
	DecimalsA        int        `json:"decimals_a"`
	DecimalsB        int        `json:"decimals_b"`
	DecimalsLP       int        `json:"decimals_lp"`
	TotalShares      Z          `json:"total_shares"`
	Shares           Z          `json:"shares"`
	ValueA           Z          `json:"value_a"`
	ValueB           Z          `json:"value_b"`
	SharesMinted     Z          `json:"shares_minted"`
	SharesBurned     Z          `json:"shares_burned"`
	SharesSent       Z          `json:"shares_sent"`
	SharesReceived   Z          `json:"shares_received"`
	DepositedA       Z          `json:"deposited_a"`
	DepositedB       Z          `json:"deposited_b"`
	WithdrawnA       Z          `json:"withdrawn_a"`
	WithdrawnB       Z          `json:"withdrawn_b"`
	IsClosed         bool       `json:"is_closed"`
	OpenBlock        int64      `json:"open_block"`
	OpenTime         time.Time  `json:"open_time"`
	CloseBlock       int64      `json:"close_block,omitempty"`
	CloseTime        *time.Time `json:"close_time,omitempty"`
	PositionValueUSD float64    `json:"position_value_usd,string"`
	ShareValueUSD    float64    `json:"share_value_usd,string"`
	OpenValueUSD     float64    `json:"open_value_usd,string"`
	CloseValueUSD    float64    `json:"close_value_usd,string"`
	FeeIncomeUSD     float64    `json:"fee_income_usd,string"`
	FeeIncomeBps     float64    `json:"fee_income_bps,string"`
	ProfitLossUSD    float64    `json:"pnl_usd,string"`
	ProfitLossBps    float64    `json:"pnl_bps,string"`
}

func (c *dexClient) ListPositions(ctx context.Context, params Query) ([]*DexPosition, error) {
	list := make([]*DexPosition, 0)
	u := params.WithPath("/v1/dex/positions").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *dexClient) ListPoolPositions(ctx context.Context, addr PoolAddress, params Query) ([]*DexPosition, error) {
	list := make([]*DexPosition, 0)
	u := params.WithPath(fmt.Sprintf("/v1/dex/%s/positions", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
