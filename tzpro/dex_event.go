// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"context"
	"fmt"
	"strconv"
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

type DexEventParams struct {
	Params
}

func NewDexEventParams() DexEventParams {
	return DexEventParams{NewParams()}
}

func (p DexEventParams) WithLimit(v uint) DexEventParams {
	p.Query.Set("limit", strconv.Itoa(int(v)))
	return p
}

func (p DexEventParams) WithOffset(v uint) DexEventParams {
	p.Query.Set("offset", strconv.Itoa(int(v)))
	return p
}

func (p DexEventParams) WithCursor(v uint64) DexEventParams {
	p.Query.Set("cursor", strconv.FormatUint(v, 10))
	return p
}

func (p DexEventParams) WithOrder(o OrderType) DexEventParams {
	p.Query.Set("order", string(o))
	return p
}

func (p DexEventParams) WithDesc() DexEventParams {
	p.Query.Set("order", string(OrderDesc))
	return p
}

func (p DexEventParams) WithAsc() DexEventParams {
	p.Query.Set("order", string(OrderAsc))
	return p
}

func (p DexEventParams) WithPool(addr tezos.Token) DexEventParams {
	p.Query.Set("pool", addr.String())
	return p
}

func (p DexEventParams) WithType(s string) DexEventParams {
	p.Query.Set("event_type", s)
	return p
}

func (p DexEventParams) WithSigner(c tezos.Address) DexEventParams {
	p.Query.Set("signer", c.String())
	return p
}

func (p DexEventParams) WithSender(c tezos.Address) DexEventParams {
	p.Query.Set("sender", c.String())
	return p
}

func (p DexEventParams) WithReceiver(c tezos.Address) DexEventParams {
	p.Query.Set("receiver", c.String())
	return p
}

func (p DexEventParams) WithTxHash(h tezos.OpHash) DexEventParams {
	p.Query.Set("tx_hash", h.String())
	return p
}

func (p DexEventParams) WithBlock(b int64) DexEventParams {
	p.Query.Set("block", ToString(b))
	return p
}

func (p DexEventParams) WithBlockEx(mode FilterMode, args ...any) DexEventParams {
	p.Query.Set("block."+string(mode), ToString(args))
	return p
}

func (p DexEventParams) WithTimeSince(t time.Time) DexEventParams {
	p.Query.Set("time.gt", t.Format(time.RFC3339))
	return p
}

func (p DexEventParams) WithTimeRange(from, to time.Time) DexEventParams {
	p.Query.Set("time.rg", from.Format(time.RFC3339)+","+to.Format(time.RFC3339))
	return p
}

func (c *Client) ListDexEvents(ctx context.Context, params DexEventParams) ([]*DexEvent, error) {
	list := make([]*DexEvent, 0)
	u := params.AppendQuery("/v1/dex/events")
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListDexPoolEvents(ctx context.Context, addr tezos.Address, id int, params DexEventParams) ([]*DexEvent, error) {
	list := make([]*DexEvent, 0)
	u := params.AppendQuery(fmt.Sprintf("/v1/dex/pools/%s_%d/events", addr, id))
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
