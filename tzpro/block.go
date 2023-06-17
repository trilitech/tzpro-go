// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package tzpro

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
)

type Block struct {
	RowId            uint64                 `json:"row_id"`
	Hash             tezos.BlockHash        `json:"hash"`
	ParentHash       *tezos.BlockHash       `json:"predecessor,omitempty"`
	FollowerHash     *tezos.BlockHash       `json:"successor,omitempty"  tzpro:"-"`
	Timestamp        time.Time              `json:"time"`
	Height           int64                  `json:"height"`
	Cycle            int64                  `json:"cycle"`
	IsCycleSnapshot  bool                   `json:"is_cycle_snapshot"`
	Solvetime        int                    `json:"solvetime"`
	Version          int                    `json:"version"`
	Round            int                    `json:"round"`
	Nonce            string                 `json:"nonce"`
	VotingPeriodKind tezos.VotingPeriodKind `json:"voting_period_kind"`
	BakerId          uint64                 `json:"baker_id"`
	Baker            tezos.Address          `json:"baker"`
	ProposerId       uint64                 `json:"proposer_id"`
	Proposer         tezos.Address          `json:"proposer"`
	NSlotsEndorsed   int                    `json:"n_endorsed_slots"`
	NOpsApplied      int                    `json:"n_ops_applied"`
	NOpsFailed       int                    `json:"n_ops_failed"`
	NContractCalls   int                    `json:"n_calls"`
	NRollupCalls     int                    `json:"n_rollup_calls"`
	NEvents          int                    `json:"n_events"`
	NTx              int                    `json:"n_tx"`
	NTickets         int                    `json:"n_tickets"`
	Volume           float64                `json:"volume"`
	Fee              float64                `json:"fee"`
	Reward           float64                `json:"reward"`
	Deposit          float64                `json:"deposit"`
	ActivatedSupply  float64                `json:"activated_supply"`
	MintedSupply     float64                `json:"minted_supply"`
	BurnedSupply     float64                `json:"burned_supply"`
	SeenAccounts     int                    `json:"n_accounts"`
	NewAccounts      int                    `json:"n_new_accounts"`
	NewContracts     int                    `json:"n_new_contracts"`
	ClearedAccounts  int                    `json:"n_cleared_accounts"`
	FundedAccounts   int                    `json:"n_funded_accounts"`
	GasLimit         int64                  `json:"gas_limit"`
	GasUsed          int64                  `json:"gas_used"`
	StoragePaid      int64                  `json:"storage_paid"`
	PctAccountReuse  float64                `json:"pct_account_reuse"`
	LbEscapeVote     string                 `json:"lb_esc_vote"`
	LbEscapeEma      int64                  `json:"lb_esc_ema"`
	Protocol         tezos.ProtocolHash     `json:"protocol"`
	ProposerKeyId    uint64                 `json:"proposer_consensus_key_id"`
	BakerKeyId       uint64                 `json:"baker_consensus_key_id"`
	ProposerKey      string                 `json:"proposer_consensus_key"`
	BakerKey         string                 `json:"baker_consensus_key"`
	Metadata         map[string]Metadata    `json:"metadata,omitempty"  tzpro:"-"`
	Rights           []Right                `json:"rights,omitempty"    tzpro:"-"`
	Ops              []*Op                  `json:"-"`
}

type Head struct {
	Hash        tezos.BlockHash `json:"hash"`
	ParentHash  tezos.BlockHash `json:"predecessor"`
	Height      int64           `json:"height"`
	Cycle       int64           `json:"cycle"`
	Timestamp   time.Time       `json:"time"`
	Baker       tezos.Address   `json:"baker"`
	Proposer    tezos.Address   `json:"proposer"`
	Round       int             `json:"round"`
	Nonce       string          `json:"nonce"`
	NOpsApplied int             `json:"n_ops_applied"`
	NOpsFailed  int             `json:"n_ops_failed"`
	Volume      float64         `json:"volume"`
	Fee         float64         `json:"fee"`
	Reward      float64         `json:"reward"`
	GasUsed     int64           `json:"gas_used"`
}

type BlockId struct {
	Height int64
	Hash   tezos.BlockHash
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
		Hash:   b.Hash.Clone(),
		Time:   b.Timestamp,
	}
}

func (b *Block) Head() *Head {
	var ph tezos.BlockHash
	if b.ParentHash != nil {
		ph = b.ParentHash.Clone()
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

type BlockList struct {
	Rows    []*Block
	columns []string
}

func (l BlockList) Len() int {
	return len(l.Rows)
}

func (l BlockList) Cursor() uint64 {
	if len(l.Rows) == 0 {
		return 0
	}
	return l.Rows[len(l.Rows)-1].RowId
}

func (l *BlockList) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, null) {
		return nil
	}
	if data[0] != '[' {
		return fmt.Errorf("BlockList: expected JSON array")
	}
	return DecodeSlice(data, l.columns, &l.Rows)
}

type BlockQuery struct {
	tableQuery
}

func (c *Client) NewBlockQuery() BlockQuery {
	return BlockQuery{c.newTableQuery("block", &Block{})}
}

func (q BlockQuery) Run(ctx context.Context) (*BlockList, error) {
	result := &BlockList{
		columns: q.Columns,
	}
	if err := q.client.QueryTable(ctx, &q.tableQuery, result); err != nil {
		return nil, err
	}
	return result, nil
}

type BlockParams = Params[Block]

func NewBlockParams() BlockParams {
	return BlockParams{
		Query: make(map[string][]string),
	}
}

func (c *Client) GetBlock(ctx context.Context, hash tezos.BlockHash, params BlockParams) (*Block, error) {
	b := &Block{}
	u := params.WithPath(fmt.Sprintf("/explorer/block/%s", hash)).Url()
	if err := c.get(ctx, u, nil, b); err != nil {
		return nil, err
	}
	return b, nil
}

func (c *Client) GetHead(ctx context.Context, params BlockParams) (*Block, error) {
	b := &Block{}
	u := params.WithPath("/explorer/block/head").Url()
	if err := c.get(ctx, u, nil, b); err != nil {
		return nil, err
	}
	return b, nil
}

func (c *Client) GetBlockHeight(ctx context.Context, height int64, params BlockParams) (*Block, error) {
	b := &Block{}
	u := params.WithPath(fmt.Sprintf("/explorer/block/%d", height)).Url()
	if err := c.get(ctx, u, nil, b); err != nil {
		return nil, err
	}
	return b, nil
}
