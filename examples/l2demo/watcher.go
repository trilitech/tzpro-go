// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package main

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go"
	"github.com/echa/log"
	"github.com/go-zeromq/zmq4"
)

type BlockWatcher struct {
	last   tzpro.BlockId
	client *tzpro.Client
	cfg    BlockWatcherConfig
	ch     chan *tzpro.Block
	addr   string
	sock   zmq4.Socket
	ctx    context.Context
	cancel context.CancelFunc
}

type BlockWatcherConfig struct {
	StartBlock   tzpro.BlockId
	MaxQueue     int
	DialTimeout  time.Duration
	RetryTimeout time.Duration
	ReadTimeout  time.Duration
}

func NewBlockWatcher(cfg BlockWatcherConfig, c *tzpro.Client) *BlockWatcher {
	return &BlockWatcher{
		last:   cfg.StartBlock,
		client: c,
		cfg:    cfg,
		ch:     make(chan *tzpro.Block, cfg.MaxQueue),
	}
}

func (w *BlockWatcher) Connect(ctx context.Context, addr string) error {
	zsub := zmq4.NewSub(ctx,
		zmq4.WithDialerRetry(w.cfg.DialTimeout),
		zmq4.WithDialerTimeout(w.cfg.RetryTimeout),
		zmq4.WithID(zmq4.SocketIdentity("xsub")),
		zmq4.WithLogger(log.Log.Logger()),
	)
	if err := zsub.Dial(addr); err != nil {
		return fmt.Errorf("Cannot connect ZMQ socket: %w", err)
	}
	if err := zsub.SetOption(zmq4.OptionSubscribe, "raw_block"); err != nil {
		return fmt.Errorf("Cannot configure ZMQ subscription: %w", err)
	}
	log.Info("ZMQ connection OK.")
	w.sock = zsub
	w.ctx, w.cancel = context.WithCancel(ctx)
	go w.recvLoop()
	return nil
}

func (w *BlockWatcher) Close() {
	if w.sock != nil {
		w.cancel()
		w.sock.Close()
		w.Flush()
		w.sock = nil
	}
}

func (w *BlockWatcher) Next() *tzpro.Block {
	return <-w.ch
}

func (w *BlockWatcher) Chan() chan *tzpro.Block {
	return w.ch
}

func (w *BlockWatcher) Flush() {
	for range w.ch {
	}
}

// simulate a socket read deadline which is unsupported by zmq4
func (w *BlockWatcher) recv() (msg zmq4.Msg, err error) {
	ctx, cancel := context.WithTimeout(w.ctx, w.cfg.ReadTimeout)
	defer cancel()
	ch := make(chan error)
	go func() {
		msg, err = w.sock.Recv()
		// close is safe here and sending err blocks for an unknown reason
		close(ch)
	}()
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case <-ch:
	}
	return
}

func (w *BlockWatcher) recvLoop() {
	// Fetch outstanding blocks
	params := tzpro.NewBlockParams()

	// close channel when done
	defer close(w.ch)

	// fetch blocks in order (reorg-free)
	for {
		var (
			b   *tzpro.Block
			err error
		)
		if w.last.Hash.IsValid() {
			b, err = w.client.GetBlock(w.ctx, w.last.Hash, params)
		} else {
			b, err = w.client.GetBlockHeight(w.ctx, w.last.Height, params)
		}
		if err != nil {
			log.Errorf("RPC: fetch block %s %d: %s", w.last.Hash, w.last.Height, err)
			select {
			case <-w.ctx.Done():
				return
			case <-time.After(5 * time.Second):
				continue
			}
		}

		// queue block, stop with context
		select {
		case <-w.ctx.Done():
			return
		case w.ch <- b:
		}

		// stop at chain tip
		if !b.FollowerHash.IsValid() {
			w.last = b.BlockId()
			break
		}

		// next block
		w.last.Height = b.Height + 1
		w.last.Hash = b.FollowerHash.Clone()
		w.last.Time = time.Time{}
	}

	log.Infof("Switching to ZMQ events")

	// Read ZMQ data
	for {
		msg, err := w.recv()
		if err != nil {
			log.Errorf("ZMQ receive: %s", err)
			return
		}
		log.Debugf("ZMQ: [%s] %d bytes", msg.Frames[0], len(msg.Frames[1]))

		if verbose {
			log.Debugf("ZMQ msg: %s", string(msg.Frames[1]))
		}

		b, err := tzpro.NewZmqMessage(msg.Frames[0], msg.Frames[1]).DecodeBlock()
		if err != nil {
			log.Errorf("ZMQ decode: %s", err)
			return
		}

		// ensure block order
		if !w.last.IsNextBlock(b) {
			log.Errorf("ZMQ invalid next block: expected parent block %d %s got %d %s",
				w.last.Height,
				w.last.Hash,
				b.Height-1,
				b.ParentHash,
			)
			return
		}
		// this block
		w.last.Height = b.Height
		w.last.Hash = b.Hash.Clone()
		w.last.Time = b.Timestamp

		// queue block, stop with context
		select {
		case <-w.ctx.Done():
			return
		case w.ch <- b:
		}
	}
}

func fetchAllBlockOps(ctx context.Context, c *tzpro.Client, hash tezos.BlockHash) ([]*tzpro.Op, error) {
	params := tzpro.NewOpParams()
	params.WithType(tzpro.FilterModeIn, "transaction,origination").
		WithCollapse().
		WithLimit(100)
	ops := make([]*tzpro.Op, 0)
	for {
		o, err := c.GetBlockOps(ctx, hash, params)
		if err != nil {
			return nil, err
		}
		if len(o) == 0 {
			break
		}
		ops = append(ops, o...)
		params = params.WithCursor(o[len(o)-1].Cursor())
	}
	return ops, nil
}
