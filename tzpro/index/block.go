// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"context"
	"fmt"
	"time"

	"github.com/trilitech/tzpro-go/internal/client"
)

type BlockAPI interface {
	GetHash(context.Context, BlockHash, Query) (*Block, error)
	GetHead(context.Context, Query) (*Block, error)
	GetHeight(context.Context, int64, Query) (*Block, error)
	ListOpsHash(context.Context, BlockHash, Query) (OpList, error)
	ListOpsHeight(context.Context, int64, Query) (OpList, error)
	NewQuery() *BlockQuery
}

func NewBlockAPI(c *client.Client) BlockAPI {
	return &blockClient{client: c}
}

type blockClient struct {
	client *client.Client
}

type Block struct {
	RowId            uint64              `json:"row_id"`
	Hash             BlockHash           `json:"hash"`
	ParentHash       *BlockHash          `json:"predecessor,omitempty"`
	FollowerHash     *BlockHash          `json:"successor,omitempty"    tzpro:"-"`
	Timestamp        time.Time           `json:"time"`
	Height           int64               `json:"height"`
	Cycle            int64               `json:"cycle"`
	IsCycleSnapshot  bool                `json:"is_cycle_snapshot"`
	Solvetime        int                 `json:"solvetime"`
	Version          int                 `json:"version"`
	Round            int                 `json:"round"`
	Nonce            string              `json:"nonce"`
	VotingPeriodKind string              `json:"voting_period_kind"`
	BakerId          uint64              `json:"baker_id"`
	Baker            Address             `json:"baker"`
	ProposerId       uint64              `json:"proposer_id"`
	Proposer         Address             `json:"proposer"`
	NSlotsEndorsed   int                 `json:"n_endorsed_slots"`
	NOpsApplied      int                 `json:"n_ops_applied"`
	NOpsFailed       int                 `json:"n_ops_failed"`
	NContractCalls   int                 `json:"n_calls"`
	NRollupCalls     int                 `json:"n_rollup_calls"`
	NEvents          int                 `json:"n_events"`
	NTx              int                 `json:"n_tx"`
	NTickets         int                 `json:"n_tickets"`
	Volume           float64             `json:"volume"`
	Fee              float64             `json:"fee"`
	Reward           float64             `json:"reward"`
	Deposit          float64             `json:"deposit"`
	ActivatedSupply  float64             `json:"activated_supply"`
	MintedSupply     float64             `json:"minted_supply"`
	BurnedSupply     float64             `json:"burned_supply"`
	SeenAccounts     int                 `json:"n_accounts"`
	NewAccounts      int                 `json:"n_new_accounts"`
	NewContracts     int                 `json:"n_new_contracts"`
	ClearedAccounts  int                 `json:"n_cleared_accounts"`
	FundedAccounts   int                 `json:"n_funded_accounts"`
	GasLimit         int64               `json:"gas_limit"`
	GasUsed          int64               `json:"gas_used"`
	StoragePaid      int64               `json:"storage_paid"`
	PctAccountReuse  float64             `json:"pct_account_reuse"`
	LbVote           string              `json:"lb_vote"`
	LbEma            int64               `json:"lb_ema"`
	AiVote           string              `json:"ai_vote"`
	AiEma            int64               `json:"ai_ema"`
	Protocol         ProtocolHash        `json:"protocol"`
	ProposerKeyId    uint64              `json:"proposer_consensus_key_id"`
	BakerKeyId       uint64              `json:"baker_consensus_key_id"`
	ProposerKey      string              `json:"proposer_consensus_key"`
	BakerKey         string              `json:"baker_consensus_key"`
	Metadata         map[string]Metadata `json:"metadata,omitempty"  tzpro:"-"`
	Rights           []Right             `json:"rights,omitempty"    tzpro:"-"`
	Ops              []*Op               `json:"-"`
}

type Head struct {
	Hash        BlockHash `json:"hash"`
	ParentHash  BlockHash `json:"predecessor"`
	Height      int64     `json:"height"`
	Cycle       int64     `json:"cycle"`
	Timestamp   time.Time `json:"time"`
	Baker       Address   `json:"baker"`
	Proposer    Address   `json:"proposer"`
	Round       int       `json:"round"`
	Nonce       string    `json:"nonce"`
	NOpsApplied int       `json:"n_ops_applied"`
	NOpsFailed  int       `json:"n_ops_failed"`
	Volume      float64   `json:"volume"`
	Fee         float64   `json:"fee"`
	Reward      float64   `json:"reward"`
	GasUsed     int64     `json:"gas_used"`
}

type BlockId struct {
	Height int64
	Hash   BlockHash
	Time   time.Time
}

func (i BlockId) IsNextBlock(b *Block) bool {
	if b == nil {
		return false
	}
	if b.Height != i.Height+1 {
		return false
	}
	if !b.ParentHash.Equal(i.Hash) {
		return false
	}
	return true
}

func (i BlockId) IsSameBlock(b *Block) bool {
	if b == nil {
		return false
	}
	if b.Height != i.Height {
		return false
	}
	if !b.Hash.Equal(i.Hash) {
		return false
	}
	return true
}

func (b *Block) BlockId() BlockId {
	return BlockId{
		Height: b.Height,
		Hash:   b.Hash,
		Time:   b.Timestamp,
	}
}

func (b *Block) Head() *Head {
	var ph BlockHash
	if b.ParentHash != nil {
		ph = *b.ParentHash
	}
	return &Head{
		Hash:        b.Hash,
		ParentHash:  ph,
		Height:      b.Height,
		Cycle:       b.Cycle,
		Timestamp:   b.Timestamp,
		Baker:       b.Baker,
		Proposer:    b.Proposer,
		Round:       b.Round,
		Nonce:       b.Nonce,
		NOpsApplied: b.NOpsApplied,
		NOpsFailed:  b.NOpsFailed,
		Volume:      b.Volume,
		Fee:         b.Fee,
		Reward:      b.Reward,
		GasUsed:     b.GasUsed,
	}
}

type BlockList []*Block

func (l BlockList) Len() int {
	return len(l)
}

func (l BlockList) Cursor() uint64 {
	if len(l) == 0 {
		return 0
	}
	return l[len(l)-1].RowId
}

type BlockQuery = client.TableQuery[*Block]

func (c blockClient) NewQuery() *BlockQuery {
	return client.NewTableQuery[*Block](c.client, "block")
}

func (c *blockClient) GetHash(ctx context.Context, hash BlockHash, params Query) (*Block, error) {
	b := &Block{}
	u := params.WithPath(fmt.Sprintf("/explorer/block/%s", hash)).Url()
	if err := c.client.Get(ctx, u, nil, b); err != nil {
		return nil, err
	}
	return b, nil
}

func (c *blockClient) GetHead(ctx context.Context, params Query) (*Block, error) {
	b := &Block{}
	u := params.WithPath("/explorer/block/head").Url()
	if err := c.client.Get(ctx, u, nil, b); err != nil {
		return nil, err
	}
	return b, nil
}

func (c *blockClient) GetHeight(ctx context.Context, height int64, params Query) (*Block, error) {
	b := &Block{}
	u := params.WithPath(fmt.Sprintf("/explorer/block/%d", height)).Url()
	if err := c.client.Get(ctx, u, nil, b); err != nil {
		return nil, err
	}
	return b, nil
}

func (c blockClient) ListOpsHash(ctx context.Context, hash BlockHash, params Query) (OpList, error) {
	ops := make(OpList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/block/%s/operations", hash)).Url()
	if err := c.client.Get(ctx, u, nil, &ops); err != nil {
		return nil, err
	}
	return ops, nil
}

func (c blockClient) ListOpsHeight(ctx context.Context, height int64, params Query) (OpList, error) {
	ops := make(OpList, 0)
	u := params.WithPath(fmt.Sprintf("/explorer/block/%d/operations", height)).Url()
	if err := c.client.Get(ctx, u, nil, &ops); err != nil {
		return nil, err
	}
	return ops, nil
}
