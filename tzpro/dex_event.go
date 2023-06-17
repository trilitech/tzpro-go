// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type DexEvent struct {
	Id         uint64       `json:"id"`
	Contract   string       `json:"contract"`
	PairId     int64        `json:"pair_id"`
	Name       string       `json:"name"`
	Entity     string       `json:"entity"`
	Pair       string       `json:"pair"`
	Type       string       `json:"event_type"`
	VolumeA    tezos.Z      `json:"vol_token_a"`
	VolumeB    tezos.Z      `json:"vol_token_b"`
	VolumeLP   tezos.Z      `json:"vol_lp"`
	DecimalsA  int          `json:"decimals_a"`
	DecimalsB  int          `json:"decimals_b"`
	DecimalsLP int          `json:"decimals_lp"`
	SupplyA    tezos.Z      `json:"supply_a"`
	SupplyB    tezos.Z      `json:"supply_b"`
	SupplyLP   tezos.Z      `json:"supply_lp"`
	Signer     string       `json:"signer"`
	Sender     string       `json:"sender"`
	Receiver   string       `json:"receiver"`
	TxHash     tezos.OpHash `json:"tx_hash"`
	TxFee      int64        `json:"tx_fee"`
	Block      int64        `json:"block"`
	Time       time.Time    `json:"time"`
}

type DexEventParams = Params[DexEvent]

func NewDexEventParams() DexEventParams {
	return DexEventParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) ListDexEvents(ctx context.Context, params DexEventParams) ([]*DexEvent, error) {
	list := make([]*DexEvent, 0)
	u := params.WithPath("/v1/dex/events").Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListDexPoolEvents(ctx context.Context, addr PoolAddress, params DexEventParams) ([]*DexEvent, error) {
	list := make([]*DexEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/dex/%s/events", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListWalletDexEvents(ctx context.Context, addr tezos.Address, params DexEventParams) ([]*DexEvent, error) {
	list := make([]*DexEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/wallets/%s/dex_events", addr)).Url()
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
