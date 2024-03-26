// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzpro-go/internal/client"
)

type AccountAPI interface {
	Get(context.Context, Address, Query) (*Account, error)
	ListOps(context.Context, Address, Query) (OpList, error)
	ListContracts(context.Context, Address, Query) (ContractList, error)
	ListTicketBalances(context.Context, Address, Query) (TicketBalanceList, error)
	ListTicketEvents(context.Context, Address, Query) (TicketEventList, error)
	NewQuery() *AccountQuery
	NewFlowQuery() *FlowQuery
}

func NewAccountAPI(c *client.Client) AccountAPI {
	return &accountClient{client: c}
}

type accountClient struct {
	client *client.Client
}

type Account struct {
	RowId              uint64              `json:"row_id"`
	Address            Address             `json:"address"`
	AddressType        AddressType         `json:"address_type"`
	Pubkey             Key                 `json:"pubkey"`
	Counter            int64               `json:"counter"`
	BakerId            uint64              `json:"baker_id,omitempty"    tzpro:"-"`
	Baker              *Address            `json:"baker,omitempty"`
	CreatorId          uint64              `json:"creator_id,omitempty"  tzpro:"-"`
	Creator            *Address            `json:"creator,omitempty"`
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
	FrozenRollupBond   float64             `json:"frozen_rollup_bond,omitempty"`
	LostRollupBond     float64             `json:"lost_rollup_bond,omitempty"`
	StakedBalance      float64             `json:"staked_balance"`
	UnstakedBalance    float64             `json:"unstaked_balance"`
	FrozenRewards      float64             `json:"frozen_rewards"         tzpro:"-"`
	LostStake          float64             `json:"lost_stake"`
	IsFunded           bool                `json:"is_funded"`
	IsActivated        bool                `json:"is_activated"`
	IsDelegated        bool                `json:"is_delegated"`
	IsStaked           bool                `json:"is_staked"`
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

type AccountList []*Account

func (l AccountList) Len() int {
	return len(l)
}

func (l AccountList) Cursor() uint64 {
	if len(l) == 0 {
		return 0
	}
	return l[len(l)-1].RowId
}

type AccountQuery = client.TableQuery[*Account]

func (c *accountClient) NewQuery() *AccountQuery {
	return client.NewTableQuery[*Account](c.client, "account")
}

func (c *accountClient) Get(ctx context.Context, addr Address, params Query) (*Account, error) {
	a := &Account{}
	u := params.WithPath(fmt.Sprintf("/explorer/account/%s", addr)).Url()
	if err := c.client.Get(ctx, u, nil, a); err != nil {
		return nil, err
	}
	return a, nil
}

func (c *accountClient) ListContracts(ctx context.Context, addr Address, params Query) (ContractList, error) {
	cc := make(ContractList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/account/%s/contracts", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &cc); err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *accountClient) ListOps(ctx context.Context, addr Address, params Query) (OpList, error) {
	ops := make(OpList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/account/%s/operations", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &ops); err != nil {
		return nil, err
	}
	return ops, nil
}

func (c *accountClient) ListTicketBalances(ctx context.Context, addr Address, params Query) (TicketBalanceList, error) {
	list := make(TicketBalanceList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/account/%s/ticket_balances", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *accountClient) ListTicketEvents(ctx context.Context, addr Address, params Query) (TicketEventList, error) {
	list := make(TicketEventList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/account/%s/ticket_events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
