// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package market

import (
	"context"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"time"

	"blockwatch.cc/tzpro-go/internal/client"
	"blockwatch.cc/tzpro-go/internal/util"
)

type Candle struct {
	Timestamp       time.Time `json:"time"`
	Open            float64   `json:"open"`
	High            float64   `json:"high"`
	Low             float64   `json:"low"`
	Close           float64   `json:"close"`
	Vwap            float64   `json:"vwap"`
	NTrades         int64     `json:"n_trades"`
	NBuy            int64     `json:"n_buy"`
	NSell           int64     `json:"n_sell"`
	VolumeBase      float64   `json:"vol_base"`
	VolumeQuote     float64   `json:"vol_quote"`
	VolumeBuyBase   float64   `json:"vol_buy_base"`
	VolumeBuyQuote  float64   `json:"vol_buy_quote"`
	VolumeSellBase  float64   `json:"vol_sell_base"`
	VolumeSellQuote float64   `json:"vol_sell_quote"`
}

type CandleList []*Candle

func (l CandleList) Len() int {
	return len(l)
}

func (l CandleList) AsOf(t time.Time) (c *Candle) {
	// when collapsing the timestamp is set to the beginning of the
	// aggregation interval (e.g. timestamp = Jun 3 means all day June 3)
	idx := sort.Search(l.Len(), func(i int) bool { return !l[i].Timestamp.Before(t) })
	if idx > 0 && idx < l.Len() {
		c = l[idx-1]
	} else {
		c = l[l.Len()-1]
	}
	return
}

type CandleQuery struct {
	Market   string
	Pair     string
	Collapse time.Duration
	Fill     client.FillMode
	Columns  []string
	From     time.Time
	To       time.Time
	Limit    int
}

func (c CandleQuery) Url() string {
	p := client.NewQuery()
	if c.Limit > 0 && p.Query.Get("limit") == "" {
		p.Query.Set("limit", strconv.Itoa(c.Limit))
	}
	if len(c.Columns) > 0 && p.Query.Get("columns") == "" {
		p.Query.Set("columns", strings.Join(c.Columns, ","))
	}
	if len(c.Fill) > 0 && p.Query.Get("fill") == "" {
		p.Query.Set("fill", string(c.Fill))
	}
	if c.Collapse > 0 && p.Query.Get("collapse") == "" {
		p.Query.Set("collapse", util.ShortDurationString(c.Collapse.String()))
	}
	if !c.From.IsZero() && p.Query.Get("start_date") == "" {
		p.Query.Set("start_date", c.From.Format(time.RFC3339))
	}
	if !c.To.IsZero() && p.Query.Get("end_date") == "" {
		p.Query.Set("end_date", c.To.Format(time.RFC3339))
	}
	return p.WithPath("/series/" + c.Market + "/" + c.Pair + "/ohlcv").Url()
}

func (c *marketClient) ListCandles(ctx context.Context, args CandleQuery) (CandleList, error) {
	var data json.RawMessage
	if err := c.client.Get(ctx, args.Url(), nil, &data); err != nil {
		return nil, err
	}
	resp := make(CandleList, 0)
	if err := client.DecodeSlice(data, args.Columns, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
