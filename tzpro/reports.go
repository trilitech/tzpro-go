// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
    "context"
    "strconv"
    "time"
)

type ReportParams struct {
    Params
}

func NewReportParams() ReportParams {
    return ReportParams{NewParams()}
}

func (p ReportParams) WithLimit(v uint) ReportParams {
    p.Query.Set("limit", strconv.Itoa(int(v)))
    return p
}

func (p ReportParams) WithOffset(v uint) ReportParams {
    p.Query.Set("offset", strconv.Itoa(int(v)))
    return p
}

func (p ReportParams) WithCursor(v uint64) ReportParams {
    p.Query.Set("cursor", strconv.FormatUint(v, 10))
    return p
}

func (p ReportParams) WithOrder(v OrderType) ReportParams {
    p.Query.Set("order", string(v))
    return p
}

func (p ReportParams) WithKind(k string) ReportParams {
    p.Query.Set("kind", k)
    return p
}

func (p ReportParams) WithType(t string) ReportParams {
    p.Query.Set("type", t)
    return p
}

func (p ReportParams) WithFrom(t time.Time) ReportParams {
    p.Query.Set("from", t.Format("2006-01-02"))
    return p
}

func (p ReportParams) WithTo(t time.Time) ReportParams {
    p.Query.Set("to", t.Format("2006-01-02"))
    return p
}

func (p ReportParams) WithRange(from, to time.Time) ReportParams {
    p.Query.Set("from", from.Format("2006-01-02"))
    p.Query.Set("to", to.Format("2006-01-02"))
    return p
}

type AgeReport struct {
    RowId      uint64    `json:"row_id"`
    Time       time.Time `json:"time"`
    Height     int64     `json:"height"`
    Year       int       `json:"year"`
    NumDormant int64     `json:"num_dormant"`
    SumDormant float64   `json:"sum_dormant"`
    columns    []string  `json:"-"`
}

func (c *Client) GetAgeReport(ctx context.Context, params ReportParams) ([]*AgeReport, error) {
    rep := make([]*AgeReport, 0)
    u := params.AppendQuery("/explorer/stats/age")
    if err := c.get(ctx, u, nil, &rep); err != nil {
        return nil, err
    }
    return rep, nil
}

type SupplyReport struct {
    Supply

    Inflation1Y             int64   `json:"inflation_1y"` // inflation
    InflationRate1Y         float64 `json:"inflation_rate_1y"`
    FutureInflationRate1Y   float64 `json:"future_inflation_rate_1y"`
    OneYearTransacting      int64   `json:"vol_tx_y1"` // 12-month HODL wave supply
    OneYearDaysDestroyed    float64 `json:"tdd_y1"`
    SixMonthTransacting     int64   `json:"vol_tx_m6"` // 6-month HODL wave supply
    SixMonthDaysDestroyed   float64 `json:"tdd_m6"`
    ThreeMonthTransacting   int64   `json:"vol_tx_m3"` // 3-month HODL wave supply
    ThreeMonthDaysDestroyed float64 `json:"tdd_m3"`
    OneMonthTransacting     int64   `json:"vol_tx_m1"` // 1-month HODL wave supply
    OneMonthDaysDestroyed   float64 `json:"tdd_m1"`
    OneWeekTransacting      int64   `json:"vol_tx_w1"` // 1-week HODL wave supply
    OneWeekDaysDestroyed    float64 `json:"tdd_w1"`
    OneDayTransacting       int64   `json:"vol_tx_d1"` // 1-day HODL wave supply
    OneDayDaysDestroyed     float64 `json:"tdd_d1"`
}

func (c *Client) GetSupplyReport(ctx context.Context, params ReportParams) ([]*SupplyReport, error) {
    rep := make([]*SupplyReport, 0)
    u := params.AppendQuery("/explorer/stats/supply")
    if err := c.get(ctx, u, nil, &rep); err != nil {
        return nil, err
    }
    return rep, nil
}

type SetsReport struct {
    RowId            uint64    `json:"row_id"`
    Time             time.Time `json:"time"`
    Height           int64     `json:"height"`
    ActiveWallets    []byte    `json:"active_wallets"`
    ActiveContracts  []byte    `json:"active_contracts"`
    GhostWallets     []byte    `json:"ghost_wallets"`
    FundedWallets    []byte    `json:"funded_wallets"`
    NewGhostWallets  []byte    `json:"new_ghost_wallets"`
    NewFundedWallets []byte    `json:"new_funded_wallets"`
}

func (c *Client) GetSetsReport(ctx context.Context, params ReportParams) ([]*SetsReport, error) {
    rep := make([]*SetsReport, 0)
    u := params.AppendQuery("/explorer/stats/sets")
    if err := c.get(ctx, u, nil, &rep); err != nil {
        return nil, err
    }
    return rep, nil
}

type ActivityReport struct {
    RowId               uint64    `json:"row_id"`
    Time                time.Time `json:"time"`
    Height              int64     `json:"height"`
    NumSeenWallets      int       `json:"num_seen_wallets"`
    NumActiveWallets    int       `json:"num_active_wallets"`
    NumActiveContracts  int       `json:"num_active_contracts"`
    NumNewWallets       int       `json:"num_new_wallets"`
    NumNewContracts     int       `json:"num_new_contracts"`
    NumNewFundedWallets int       `json:"num_new_funded_wallets"`
    NumNewGhostWallets  int       `json:"num_new_ghost_wallets"`
    NumClearedWallets   int       `json:"num_cleared_wallets"`
    SumVolume           float64   `json:"sum_volume"`
    NumTx               int       `json:"num_tx"`
    NumContractCalls    int       `json:"num_contract_calls"`
    SumVolumeTop1       float64   `json:"vol_top1"`
    SumVolumeTop10      float64   `json:"vol_top10"`
    SumVolumeTop100     float64   `json:"vol_top100"`
    SumVolumeTop1k      float64   `json:"vol_top1k"`
    SumVolumeTop10k     float64   `json:"vol_top10k"`
    SumVolumeTop100k    float64   `json:"vol_top100k"`
    TrafficTop1         int       `json:"num_tx_top1"`
    TrafficTop10        int       `json:"num_tx_top10"`
    TrafficTop100       int       `json:"num_tx_top100"`
    TrafficTop1k        int       `json:"num_tx_top1k"`
    TrafficTop10k       int       `json:"num_tx_top10k"`
    TrafficTop100k      int       `json:"num_tx_top100k"`
}

func (c *Client) GetActivityReport(ctx context.Context, params ReportParams) ([]*ActivityReport, error) {
    rep := make([]*ActivityReport, 0)
    u := params.AppendQuery("/explorer/stats/activity")
    if err := c.get(ctx, u, nil, &rep); err != nil {
        return nil, err
    }
    return rep, nil
}

type BalanceReport struct {
    RowId              uint64    `json:"row_id"`
    Time               time.Time `json:"time"`
    Height             int64     `json:"height"`
    NumFundedTotal     int       `json:"num_funded_total"`
    NumFundedWallets   int       `json:"num_funded_wallets"`
    NumFundedContracts int       `json:"num_funded_contracts"`
    SumFundedWallets   float64   `json:"sum_wallets"`
    SumFundedContracts float64   `json:"sum_contracts"`
    BalanceTop1        float64   `json:"balance_top1"`
    BalanceTop10       float64   `json:"balance_top10"`
    BalanceTop100      float64   `json:"balance_top100"`
    BalanceTop1k       float64   `json:"balance_top1k"`
    BalanceTop10k      float64   `json:"balance_top10k"`
    BalanceTop100k     float64   `json:"balance_top100k"`
    NumBalanceHist1E0  int       `json:"hist_num_bal_1e0"`
    NumBalanceHist1E1  int       `json:"hist_num_bal_1e1"`
    NumBalanceHist1E2  int       `json:"hist_num_bal_1e2"`
    NumBalanceHist1E3  int       `json:"hist_num_bal_1e3"`
    NumBalanceHist1E4  int       `json:"hist_num_bal_1e4"`
    NumBalanceHist1E5  int       `json:"hist_num_bal_1e5"`
    NumBalanceHist1E6  int       `json:"hist_num_bal_1e6"`
    NumBalanceHist1E7  int       `json:"hist_num_bal_1e7"`
    NumBalanceHist1E8  int       `json:"hist_num_bal_1e8"`
    NumBalanceHist1E9  int       `json:"hist_num_bal_1e9"`
    NumBalanceHist1E10 int       `json:"hist_num_bal_1e10"`
    NumBalanceHist1E11 int       `json:"hist_num_bal_1e11"`
    NumBalanceHist1E12 int       `json:"hist_num_bal_1e12"`
    NumBalanceHist1E13 int       `json:"hist_num_bal_1e13"`
    NumBalanceHist1E14 int       `json:"hist_num_bal_1e14"`
    NumBalanceHist1E15 int       `json:"hist_num_bal_1e15"`
    NumBalanceHist1E16 int       `json:"hist_num_bal_1e16"`
    NumBalanceHist1E17 int       `json:"hist_num_bal_1e17"`
    NumBalanceHist1E18 int       `json:"hist_num_bal_1e18"`
    NumBalanceHist1E19 int       `json:"hist_num_bal_1e19"`
    SumBalanceHist1E0  float64   `json:"hist_sum_bal_1e0"`
    SumBalanceHist1E1  float64   `json:"hist_sum_bal_1e1"`
    SumBalanceHist1E2  float64   `json:"hist_sum_bal_1e2"`
    SumBalanceHist1E3  float64   `json:"hist_sum_bal_1e3"`
    SumBalanceHist1E4  float64   `json:"hist_sum_bal_1e4"`
    SumBalanceHist1E5  float64   `json:"hist_sum_bal_1e5"`
    SumBalanceHist1E6  float64   `json:"hist_sum_bal_1e6"`
    SumBalanceHist1E7  float64   `json:"hist_sum_bal_1e7"`
    SumBalanceHist1E8  float64   `json:"hist_sum_bal_1e8"`
    SumBalanceHist1E9  float64   `json:"hist_sum_bal_1e9"`
    SumBalanceHist1E10 float64   `json:"hist_sum_bal_1e10"`
    SumBalanceHist1E11 float64   `json:"hist_sum_bal_1e11"`
    SumBalanceHist1E12 float64   `json:"hist_sum_bal_1e12"`
    SumBalanceHist1E13 float64   `json:"hist_sum_bal_1e13"`
    SumBalanceHist1E14 float64   `json:"hist_sum_bal_1e14"`
    SumBalanceHist1E15 float64   `json:"hist_sum_bal_1e15"`
    SumBalanceHist1E16 float64   `json:"hist_sum_bal_1e16"`
    SumBalanceHist1E17 float64   `json:"hist_sum_bal_1e17"`
    SumBalanceHist1E18 float64   `json:"hist_sum_bal_1e18"`
    GiniFunded         float64   `json:"gini_funded"`
    GiniOneTz          float64   `json:"gini_onetez"`
    GiniBakers         float64   `json:"gini_bakers"`
}

func (c *Client) GetBalanceReport(ctx context.Context, params ReportParams) ([]*BalanceReport, error) {
    rep := make([]*BalanceReport, 0)
    u := params.AppendQuery("/explorer/stats/balance")
    if err := c.get(ctx, u, nil, &rep); err != nil {
        return nil, err
    }
    return rep, nil
}

// fee, gas, vol, tdd, add
type OpReport struct {
    RowId  uint64    `json:"row_id"`
    Time   time.Time `json:"time"`
    Height int64     `json:"height"`
    Kind   string    `json:"kind"`
    Type   OpType    `json:"type"`
    N      int       `json:"count"`
    Sum    float64   `json:"sum"`
    Min    float64   `json:"min"`
    Max    float64   `json:"max"`
    Mean   float64   `json:"mean"`
    Median float64   `json:"median"`
}

func (c *Client) GetOpReport(ctx context.Context, params ReportParams) ([]*OpReport, error) {
    rep := make([]*OpReport, 0)
    u := params.AppendQuery("/explorer/stats/op")
    if err := c.get(ctx, u, nil, &rep); err != nil {
        return nil, err
    }
    return rep, nil
}
