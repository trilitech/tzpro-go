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
	Contract         string     `json:"contract"`
	PairId           int64      `json:"pair_id"`
	Name             string     `json:"name"`
	Entity           string     `json:"entity"`
	Pair             string     `json:"pair"`
	Owner            string     `json:"owner"`
	DecimalsA        int        `json:"decimals_a"`
	DecimalsB        int        `json:"decimals_b"`
	DecimalsLP       int        `json:"decimals_lp"`
	TotalShares      string     `json:"total_shares"`
	Shares           string     `json:"shares"`
	ValueA           string     `json:"value_a"`
	ValueB           string     `json:"value_b"`
	SharesMinted     string     `json:"shares_minted"`
	SharesBurned     string     `json:"shares_burned"`
	SharesSent       string     `json:"shares_sent"`
	SharesReceived   string     `json:"shares_received"`
	DepositedA       string     `json:"deposited_a"`
	DepositedB       string     `json:"deposited_b"`
	WithdrawnA       string     `json:"withdrawn_a"`
	WithdrawnB       string     `json:"withdrawn_b"`
	IsClosed         bool       `json:"is_closed"`
	OpenBlock        int64      `json:"open_block"`
	OpenTime         time.Time  `json:"open_time"`
	CloseBlock       int64      `json:"close_block,omitempty"`
	CloseTime        *time.Time `json:"close_time,omitempty"`
	PositionValueUSD string     `json:"position_value_usd"`
	ShareValueUSD    string     `json:"share_value_usd"`
	ProfitLossUSD    string     `json:"pnl_usd"`
	ProfitLossBps    float64    `json:"pnl_bps,string"`
}

func (c *dexClient) ListPositions(ctx context.Context, params Params) ([]*DexPosition, error) {
	list := make([]*DexPosition, 0)
	u := params.WithPath("/v1/dex/positions").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *dexClient) ListPoolPositions(ctx context.Context, addr PoolAddress, params Params) ([]*DexPosition, error) {
	list := make([]*DexPosition, 0)
	u := params.WithPath(fmt.Sprintf("/v1/dex/%s/positions", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
