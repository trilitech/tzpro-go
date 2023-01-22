// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
    "context"
    "fmt"
    // "strconv"
    "time"

    "blockwatch.cc/tzgo/tezos"
)

type DexTicker struct {
    Pair             string    `json:"pair"`
    Pool             string    `json:"pool"`
    Name             string    `json:"name"`
    Entity           string    `json:"entity"`
    PriceChange      string    `json:"price_change"`
    PriceChangeBps   string    `json:"price_change_bps"`
    AskPrice         string    `json:"ask_price"`
    LiquidityUSD     string    `json:"liquidity_usd"`
    WeightedAvgPrice string    `json:"weighted_avg_price"`
    LastPrice        string    `json:"last_price"`
    LastQty          string    `json:"last_qty"`
    LastTradeTime    string    `json:"last_trade_time"`
    BaseVolume       string    `json:"base_volume"`
    QuoteVolume      string    `json:"quote_volume"`
    OpenPrice        string    `json:"open_price"`
    HighPrice        string    `json:"high_price"`
    LowPrice         string    `json:"low_price"`
    OpenTime         time.Time `json:"open_time"`
    CloseTime        time.Time `json:"close_time"`
    NumTrades        int       `json:"num_trades"`
}

type DexTickerParams struct {
    Params
}

func NewDexTickerParams() DexTickerParams {
    return DexTickerParams{NewParams()}
}

// func (p DexTickerParams) WithLimit(v uint) DexTickerParams {
//     p.Query.Set("limit", strconv.Itoa(int(v)))
//     return p
// }

// func (p DexTickerParams) WithOffset(v uint) DexTickerParams {
//     p.Query.Set("offset", strconv.Itoa(int(v)))
//     return p
// }

// func (p DexTickerParams) WithCursor(v uint64) DexTickerParams {
//     p.Query.Set("cursor", strconv.FormatUint(v, 10))
//     return p
// }

// func (p DexTickerParams) WithOrder(o OrderType) DexTickerParams {
//     p.Query.Set("order", string(o))
//     return p
// }

// func (p DexTickerParams) WithDesc() DexTickerParams {
//     p.Query.Set("order", string(OrderDesc))
//     return p
// }

// func (p DexTickerParams) WithAsc() DexTickerParams {
//     p.Query.Set("order", string(OrderAsc))
//     return p
// }

func (c *Client) GetDexTicker(ctx context.Context, addr tezos.Address, id int) (*DexTicker, error) {
    tick := &DexTicker{}
    u := fmt.Sprintf("/v1/dex/pools/%s_%d/ticker", addr, id)
    if err := c.get(ctx, u, nil, tick); err != nil {
        return nil, err
    }
    return tick, nil
}

func (c *Client) ListDexTickers(ctx context.Context, params DexTickerParams) ([]*DexTicker, error) {
    list := make([]*DexTicker, 0)
    u := params.AppendQuery("/v1/dex/tickers")
    if err := c.get(ctx, u, nil, &list); err != nil {
        return nil, err
    }
    return list, nil
}
