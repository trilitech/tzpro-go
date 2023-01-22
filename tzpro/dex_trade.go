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

type DexTrade struct {
	Id             uint64       `json:"id"`
	Contract       string       `json:"contract"`
	PairId         int64        `json:"pair_id"`
	Name           string       `json:"name"`
	Entity         string       `json:"entity"`
	Pair           string       `json:"pair"`
	Counter        int64        `json:"counter"`
	Side           string       `json:"side"`
	BaseVolume     tezos.Z      `json:"base_volume"`
	BaseSymbol     string       `json:"base_symbol"`
	BaseDecimals   int          `json:"base_decimals"`
	QuoteVolume    tezos.Z      `json:"quote_volume"`
	QuoteSymbol    string       `json:"quote_symbol"`
	QuoteDecimals  int          `json:"quote_decimals"`
	LpFee          tezos.Z      `json:"lp_fee"`
	LpFeeBps       float64      `json:"lp_fee_bps,string"`
	LpFeeSymbol    string       `json:"lp_fee_symbol"`
	LpFeeDecimals  int          `json:"lp_fee_decimals"`
	DevFee         tezos.Z      `json:"dev_fee"`
	DevFeeBps      float64      `json:"dev_fee_bps,string"`
	DevFeeSymbol   string       `json:"dev_fee_symbol"`
	DevFeeDecimals int          `json:"dev_fee_decimals"`
	RefFee         tezos.Z      `json:"referral_fee"`
	RefFeeBps      float64      `json:"referral_fee_bps,string"`
	RefFeeSymbol   string       `json:"referral_fee_symbol"`
	RefFeeDecimals int          `json:"referral_fee_decimals"`
	IncFee         tezos.Z      `json:"incentive_fee"`
	IncFeeBps      float64      `json:"incentive_fee_bps,string"`
	IncFeeSymbol   string       `json:"incentive_fee_symbol"`
	IncFeeDecimals int          `json:"incentive_fee_decimals"`
	Burn           tezos.Z      `json:"burn"`
	BurnBps        float64      `json:"burn_bps,string"`
	BurnSymbol     string       `json:"burn_symbol"`
	BurnDecimals   int          `json:"burn_decimals"`
	PriceSymbol    string       `json:"price_symbol"`
	PriceDecimals  int          `json:"price_decimals"`
	PriceNet       tezos.Z      `json:"price_net"`               // including fees
	PriceGross     tezos.Z      `json:"price_gross"`             // excluding fees
	PriceBefore    tezos.Z      `json:"price_before"`            // marginal price before execution
	PriceAfter     tezos.Z      `json:"price_after"`             // marginal price after execution
	Delta          float64      `json:"price_delta_bps,string"`  // price delta (in basispoints) between mid price and execution price
	Impact         float64      `json:"price_impact_bps,string"` // price impact (in basispoints) between mid price and next mid price
	Signer         string       `json:"signer"`
	Sender         string       `json:"sender"`
	Receiver       string       `json:"receiver"`
	IsWash         bool         `json:"is_wash_trade"`
	TxHash         tezos.OpHash `json:"tx_hash"`
	TxFee          int64        `json:"tx_fee,string"`
	TxFeeSymbol    string       `json:"tx_fee_symbol"`
	TxFeeDecimals  int          `json:"tx_fee_decimals"`
	Block          int64        `json:"block"`
	Time           time.Time    `json:"time"`
	PriceUSD       tezos.Z      `json:"price_usd"`
	FeesUSD        tezos.Z      `json:"fees_usd"`
	VolumeUSD      tezos.Z      `json:"volume_usd"`
}

type DexTradeParams struct {
	Params
}

func NewDexTradeParams() DexTradeParams {
	return DexTradeParams{NewParams()}
}

func (p DexTradeParams) WithLimit(v uint) DexTradeParams {
	p.Query.Set("limit", strconv.Itoa(int(v)))
	return p
}

func (p DexTradeParams) WithOffset(v uint) DexTradeParams {
	p.Query.Set("offset", strconv.Itoa(int(v)))
	return p
}

func (p DexTradeParams) WithCursor(v uint64) DexTradeParams {
	p.Query.Set("cursor", strconv.FormatUint(v, 10))
	return p
}

func (p DexTradeParams) WithOrder(o OrderType) DexTradeParams {
	p.Query.Set("order", string(o))
	return p
}

func (p DexTradeParams) WithDesc() DexTradeParams {
	p.Query.Set("order", string(OrderDesc))
	return p
}

func (p DexTradeParams) WithAsc() DexTradeParams {
	p.Query.Set("order", string(OrderAsc))
	return p
}

func (p DexTradeParams) WithSide(s string) DexTradeParams {
	p.Query.Set("side", s)
	return p
}

func (p DexTradeParams) WithCounter(c int64) DexTradeParams {
	p.Query.Set("counter", strconv.FormatInt(c, 10))
	return p
}

func (p DexTradeParams) WithSigner(c tezos.Address) DexTradeParams {
	p.Query.Set("signer", c.String())
	return p
}

func (p DexTradeParams) WithSender(c tezos.Address) DexTradeParams {
	p.Query.Set("sender", c.String())
	return p
}

func (p DexTradeParams) WithReceiver(c tezos.Address) DexTradeParams {
	p.Query.Set("receiver", c.String())
	return p
}

func (p DexTradeParams) WithTxHash(h tezos.OpHash) DexTradeParams {
	p.Query.Set("tx_hash", h.String())
	return p
}

func (p DexTradeParams) WithWashTrades(b bool) DexTradeParams {
	p.Query.Set("is_wash_trade", strconv.FormatBool(b))
	return p
}

func (p DexTradeParams) WithTimeSince(t time.Time) DexTradeParams {
	p.Query.Set("time.gt", t.Format(time.RFC3339))
	return p
}

func (p DexTradeParams) WithTimeRange(from, to time.Time) DexTradeParams {
	p.Query.Set("time.rg", from.Format(time.RFC3339)+","+to.Format(time.RFC3339))
	return p
}

func (c *Client) ListDexTrades(ctx context.Context, params DexTradeParams) ([]*DexTrade, error) {
	list := make([]*DexTrade, 0)
	u := params.AppendQuery("/v1/dex/trades")
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) ListDexPoolTrades(ctx context.Context, addr tezos.Address, id int, params DexTradeParams) ([]*DexTrade, error) {
	list := make([]*DexTrade, 0)
	u := params.AppendQuery(fmt.Sprintf("/v1/dex/pools/%s_%d/trades", addr, id))
	if err := c.get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
