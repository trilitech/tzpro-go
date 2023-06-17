// Copyright (c) 2020-2022 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type Account struct {
	RowId              uint64              `json:"row_id"`
	Address            tezos.Address       `json:"address"`
	AddressType        tezos.AddressType   `json:"address_type"`
	Pubkey             tezos.Key           `json:"pubkey"`
	Counter            int64               `json:"counter"`
	BakerId            uint64              `json:"baker_id,omitempty"    tzpro:"-"`
	Baker              *tezos.Address      `json:"baker,omitempty"`
	CreatorId          uint64              `json:"creator_id,omitempty"  tzpro:"-"`
	Creator            *tezos.Address      `json:"creator,omitempty"`
	FirstIn            int64               `json:"first_in"`
	FirstOut           int64               `json:"first_out"`
	FirstSeen          int64               `json:"first_seen"`
	LastIn             int64               `json:"last_in"`
	LastOut            int64               `json:"last_out"`
	LastSeen           int64               `json:"last_seen"`
	FirstSeenTime      time.Time           `json:"first_seen_time"`
	LastSeenTime       time.Time           `json:"last_seen_time"`
	FirstInTime        time.Time           `json:"first_in_time"`
	LastInTime         time.Time           `json:"last_in_time"`
	FirstOutTime       time.Time           `json:"first_out_time"`
	LastOutTime        time.Time           `json:"last_out_time"`
	DelegatedSince     int64               `json:"delegated_since"`
	DelegatedSinceTime time.Time           `json:"delegated_since_time"`
	TotalReceived      float64             `json:"total_received"`
	TotalSent          float64             `json:"total_sent"`
	TotalBurned        float64             `json:"total_burned"`
	TotalFeesPaid      float64             `json:"total_fees_paid"`
	TotalFeesUsed      float64             `json:"total_fees_used"`
	UnclaimedBalance   float64             `json:"unclaimed_balance,omitempty"`
	SpendableBalance   float64             `json:"spendable_balance"`
	FrozenBond         float64             `json:"frozen_bond"`
	LostBond           float64             `json:"lost_bond"`
	IsFunded           bool                `json:"is_funded"`
	IsActivated        bool                `json:"is_activated"`
	IsDelegated        bool                `json:"is_delegated"`
	IsRevealed         bool                `json:"is_revealed"`
	IsBaker            bool                `json:"is_baker"`
	IsContract         bool                `json:"is_contract"`
	NTxSuccess         int                 `json:"n_tx_success"`
	NTxFailed          int                 `json:"n_tx_failed"`
	NTxOut             int                 `json:"n_tx_out"`
	NTxIn              int                 `json:"n_tx_in"`
	LifetimeRewards    float64             `json:"lifetime_rewards,omitempty" tzpro:"-"`
	PendingRewards     float64             `json:"pending_rewards,omitempty"  tzpro:"-"`
	Metadata           map[string]Metadata `json:"metadata,omitempty"         tzpro:"-"`
}

type AccountList struct {
	Rows    []*Account
	columns []string
}

func (l AccountList) Len() int {
	return len(l.Rows)
}

func (l AccountList) Cursor() uint64 {
	if len(l.Rows) == 0 {
		return 0
	}
	return l.Rows[len(l.Rows)-1].RowId
}

func (l *AccountList) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, null) {
		return nil
	}
	if data[0] != '[' {
		return fmt.Errorf("AccountList: expected JSON array")
	}
	return DecodeSlice(data, l.columns, &l.Rows)
}

type AccountParams = Params[Account]

func NewAccountParams() AccountParams {
	return AccountParams{
		Query: make(map[string][]string),
	}
}

type AccountQuery struct {
	tableQuery
}

func (c *Client) NewAccountQuery() AccountQuery {
	return AccountQuery{c.newTableQuery("account", &Account{})}
}

func (q AccountQuery) Run(ctx context.Context) (*AccountList, error) {
	result := &AccountList{
		columns: q.Columns,
	}
	if err := q.client.QueryTable(ctx, &q.tableQuery, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) QueryAccounts(ctx context.Context, filter FilterList, cols []string) (*AccountList, error) {
	q := c.NewAccountQuery()
	if len(cols) > 0 {
		q.Columns = cols
	}
	if len(filter) > 0 {
		q.Filter = filter
	}
	return q.Run(ctx)
}

func (c *Client) GetAccount(ctx context.Context, addr tezos.Address, params AccountParams) (*Account, error) {
	a := &Account{}
	u := params.WithPath(fmt.Sprintf("/explorer/account/%s", addr)).Url()
	if err := c.get(ctx, u, nil, a); err != nil {
		return nil, err
	}
	return a, nil
}

func (c *Client) GetAccountContracts(ctx context.Context, addr tezos.Address, params AccountParams) ([]*Account, error) {
	cc := make([]*Account, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/account/%s/contracts", addr)).Url()
	if err := c.get(ctx, u, nil, &cc); err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *Client) GetAccountOps(ctx context.Context, addr tezos.Address, params OpParams) ([]*Op, error) {
	ops := make([]*Op, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/account/%s/operations", addr)).Url()
	if err := c.get(ctx, u, nil, &ops); err != nil {
		return nil, err
	}
	return ops, nil
}
