// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"
	"fmt"
	"time"
)

type DexTrade struct {
	Id             uint64    `json:"id"`
	Contract       string    `json:"contract"`
	PairId         int       `json:"pair_id"`
	Name           string    `json:"name"`
	Entity         string    `json:"entity"`
	Pair           string    `json:"pair"`
	Counter        int64     `json:"counter"`
	Side           string    `json:"side"`
	BaseVolume     Z         `json:"base_volume"`
	BaseSymbol     string    `json:"base_symbol"`
	BaseDecimals   int       `json:"base_decimals"`
	QuoteVolume    Z         `json:"quote_volume"`
	QuoteSymbol    string    `json:"quote_symbol"`
	QuoteDecimals  int       `json:"quote_decimals"`
	LpFee          Z         `json:"lp_fee"`
	LpFeeBps       float64   `json:"lp_fee_bps,string"`
	LpFeeSymbol    string    `json:"lp_fee_symbol"`
	LpFeeDecimals  int       `json:"lp_fee_decimals"`
	DevFee         Z         `json:"dev_fee"`
	DevFeeBps      float64   `json:"dev_fee_bps,string"`
	DevFeeSymbol   string    `json:"dev_fee_symbol"`
	DevFeeDecimals int       `json:"dev_fee_decimals"`
	RefFee         Z         `json:"referral_fee"`
	RefFeeBps      float64   `json:"referral_fee_bps,string"`
	RefFeeSymbol   string    `json:"referral_fee_symbol"`
	RefFeeDecimals int       `json:"referral_fee_decimals"`
	IncFee         Z         `json:"incentive_fee"`
	IncFeeBps      float64   `json:"incentive_fee_bps,string"`
	IncFeeSymbol   string    `json:"incentive_fee_symbol"`
	IncFeeDecimals int       `json:"incentive_fee_decimals"`
	Burn           Z         `json:"burn"`
	BurnBps        float64   `json:"burn_bps,string"`
	BurnSymbol     string    `json:"burn_symbol"`
	BurnDecimals   int       `json:"burn_decimals"`
	PriceSymbol    string    `json:"price_symbol"`
	PriceDecimals  int       `json:"price_decimals"`
	PriceNet       float64   `json:"price_net,string"`        // including fees
	PriceGross     float64   `json:"price_gross,string"`      // excluding fees
	PriceBefore    float64   `json:"price_before,string"`     // marginal price before execution
	PriceAfter     float64   `json:"price_after,string"`      // marginal price after execution
	Delta          float64   `json:"price_delta_bps,string"`  // price delta (in basispoints) between mid price and execution price
	Impact         float64   `json:"price_impact_bps,string"` // price impact (in basispoints) between mid price and next mid price
	Signer         string    `json:"signer"`
	Sender         string    `json:"sender"`
	Receiver       string    `json:"receiver"`
	IsWash         bool      `json:"is_wash_trade"`
	TxHash         string    `json:"tx_hash"`
	TxFee          int64     `json:"tx_fee,string"`
	TxFeeSymbol    string    `json:"tx_fee_symbol"`
	TxFeeDecimals  int       `json:"tx_fee_decimals"`
	Block          int64     `json:"block"`
	Time           time.Time `json:"time"`
	PriceUSD       float64   `json:"price_usd,string"`
	FeesUSD        float64   `json:"fees_usd,string"`
	VolumeUSD      float64   `json:"volume_usd,string"`
}

func (c *dexClient) ListTrades(ctx context.Context, params Query) ([]*DexTrade, error) {
	list := make([]*DexTrade, 0)
	u := params.WithPath("/v1/dex/trades").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *dexClient) ListPoolTrades(ctx context.Context, addr PoolAddress, params Query) ([]*DexTrade, error) {
	list := make([]*DexTrade, 0)
	u := params.WithPath(fmt.Sprintf("/v1/dex/%s/trades", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
