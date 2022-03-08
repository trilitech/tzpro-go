// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package provider

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go"
	"github.com/echa/log"
)

func init() {
	register(NewGenericProvider())
}

type GenericProvider struct {
	address tezos.Address // contract address
	client  *tzpro.Client

	// last update
	status tzpro.BlockId

	// contract type info
	contract *tzpro.Contract
	script   *tzpro.ContractScript
}

var _ Provider = (*GenericProvider)(nil)

func NewGenericProvider() Provider {
	return &GenericProvider{}
}

func (p *GenericProvider) Name() string {
	return "generic"
}

func (p *GenericProvider) Status() tzpro.BlockId {
	return p.status
}

func (p *GenericProvider) Enabled() bool {
	return true
}

func (p *GenericProvider) Init(ctx context.Context, c *tzpro.Client, addr tezos.Address) error {
	p.address = addr
	p.client = c

	// only successful when contract is deployed
	var err error
	params := tzpro.NewContractParams().WithPrim()
	p.contract, err = p.client.GetContract(ctx, p.address, params)
	if err != nil {
		return fmt.Errorf("generic: loading contract %s: %w", p.address, err)
	}

	p.script, err = p.client.GetContractScript(ctx, p.address, params)
	if err != nil {
		return fmt.Errorf("generic: loading contract %s script: %w", p.address, err)
	}

	p.status.Height = p.contract.FirstSeen
	p.status.Time = p.contract.FirstSeenTime

	log.Infof("generic: watching contract %s from %d (%s)", p.address, p.status.Height, p.status.Time)
	return nil
}

func (p *GenericProvider) FillHistory(ctx context.Context) error {
	// fetch all successful transactions sent to the contract,
	// load parameters and storage updates for analysis
	// Note: also load the very first call which is the contract origination
	q := p.client.NewOpQuery()
	q.WithColumns("row_id", "height", "time", "hash", "block_hash", "type", "parameters", "storage", "big_map_diff").
		WithFilter(tzpro.FilterModeGt, "height", p.status.Height).
		WithFilter(tzpro.FilterModeEqual, "receiver", p.address.String()).
		WithFilter(tzpro.FilterModeIn, "type", "transaction,origination").
		WithFilter(tzpro.FilterModeEqual, "is_success", true).
		WithLimit(50_000)

	// get current chain state
	// tip, err := p.client.GetTip(ctx)
	// if err != nil {
	// 	return fmt.Errorf("generic: loading tip: %w", err)
	// }

	// track the last processed block
	last := p.status

	// fetch all contract calls sequentially
	for {
		res, err := q.Run(ctx)
		if err != nil {
			return fmt.Errorf("generic: loading contract call history: %w", err)
		}

		// stop fetching when no more calls are
		if res.Len() == 0 {
			break
		}

		for _, op := range res.Rows {
			if err := p.ConnectOp(ctx, op); err != nil {
				return fmt.Errorf("generic: process op %s [%d]: %w",
					op.Hash, op.Height, err)
			}

			// track last height (note: there may be more updates at this height
			// after the current row, so don't save just yet)
			last = op.BlockId()
		}

		// advance query cursor to last processed operation
		q.Cursor = res.Cursor()
	}

	// save crawl state
	return p.saveState(ctx, last)
}

func (p *GenericProvider) ConnectBlock(ctx context.Context, block *tzpro.Block) error {
	// ignore already processed blocks
	if block.Height <= p.status.Height {
		return nil
	}

	var (
		saveState bool
		last      tzpro.BlockId
	)
	for _, op := range block.Ops {
		// skip failed op groups
		if !op.IsSuccess {
			continue
		}

		// process updates
		for _, o := range op.Content() {
			// skip all ops that do not match our contract
			if o.Type != tzpro.OpTypeTransaction || !o.Receiver.Equal(p.address) {
				continue
			}
			if err := p.ConnectOp(ctx, o); err != nil {
				return fmt.Errorf("generic: process op %s [%d]: %w", o.Hash, o.Height, err)
			}
			last = o.BlockId()
			saveState = true
		}
	}

	// save a state snapshot on update
	if saveState {
		if err := p.saveState(ctx, last); err != nil {
			return err
		}
	}
	return nil
}

func (p *GenericProvider) DisconnectBlock(ctx context.Context, block *tzpro.Block) error {
	// ignore already processed blocks
	if block.Height > p.status.Height {
		return nil
	}

	// ops are in reverse order
	var saveState bool
	for _, op := range block.Ops {
		// skip failed op groups
		if !op.IsSuccess {
			continue
		}

		// process bigmap updates
		for _, o := range op.Content() {
			// skip all ops that do not match our contract
			if o.Type != tzpro.OpTypeTransaction || !o.Receiver.Equal(p.address) {
				continue
			}

			if err := p.DisconnectOp(ctx, o); err != nil {
				return fmt.Errorf("generic: process op %s [%d]: %w", o.Hash, o.Height, err)
			}
			saveState = true
		}
	}

	// rollback state to previous block
	if saveState && block.ParentHash != nil {
		id := tzpro.BlockId{
			Hash:   *block.ParentHash,
			Height: block.Height - 1,
			Time:   time.Time{},
		}
		if err := p.saveState(ctx, id); err != nil {
			return err
		}
	}
	return nil
}

func (p *GenericProvider) ConnectOp(ctx context.Context, op *tzpro.Op) error {
	log.Infof("Processing %s %s", op.Hash, op.Parameters.Entrypoint)
	return nil
}

func (p *GenericProvider) DisconnectOp(ctx context.Context, op *tzpro.Op) error {
	log.Infof("Rolling back %s %s", op.Hash, op.Parameters.Entrypoint)
	return nil
}

func (p *GenericProvider) saveState(ctx context.Context, id tzpro.BlockId) error {
	// TODO: store most recent processed block somewhere

	// update internal status
	p.status = id
	return nil
}
