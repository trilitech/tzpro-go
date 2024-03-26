// Copyright (c) 2020-2024 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
    "time"

    "blockwatch.cc/tzpro-go/internal/client"
)

type Flow struct {
    Id             uint64    `json:"id"`
    Height         int64     `json:"height"`
    Cycle          int64     `json:"cycle"`
    Timestamp      time.Time `json:"time"`
    OpN            int       `json:"op_n"`
    OpC            int       `json:"op_c"`
    OpI            int       `json:"op_i"`
    AccountId      uint64    `json:"account_id"`
    Account        Address   `json:"address"`
    CounterPartyId uint64    `json:"counterparty_id"`
    CounterParty   Address   `json:"counterparty"`
    Kind           string    `json:"kind"`
    Type           string    `json:"type"`
    AmountIn       float64   `json:"amount_in"`
    AmountOut      float64   `json:"amount_out"`
    IsFee          bool      `json:"is_fee"`
    IsBurned       bool      `json:"is_burned"`
    IsFrozen       bool      `json:"is_frozen"`
    IsUnfrozen     bool      `json:"is_unfrozen"`
    IsShielded     bool      `json:"is_shielded"`
    IsUnshielded   bool      `json:"is_unshielded"`
    TokenAge       int64     `json:"token_age"`
}

type FlowQuery = client.TableQuery[*Flow]

func (a accountClient) NewFlowQuery() *FlowQuery {
    return client.NewTableQuery[*Flow](a.client, "flow")
}
