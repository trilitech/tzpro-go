// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go"
	"blockwatch.cc/tzpro-go/examples/l2demo/provider"
	"github.com/echa/config"
	"github.com/echa/log"
)

func init() {
	config.SetDefault("max_queue", 100)
	config.SetDefault("zmq.dial_timeout", 2*time.Second)
	config.SetDefault("zmq.dial_retry", 5*time.Second)
	config.SetDefault("zmq.read_timeout", 90*time.Second)
}

func watchContract(ctx context.Context, c *tzpro.Client, addr tezos.Address) error {
	// init providers
	for _, p := range provider.Providers() {
		if err := p.Init(ctx, c, addr); err != nil {
			return err
		}
	}

	// watch for shutdown signals
	sctx, stop := signal.NotifyContext(ctx,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer stop()

	// close ZMQ socket
	var bw *BlockWatcher
	defer func() {
		if bw != nil {
			bw.Close()
			bw = nil
		}
	}()

	plog := log.NewProgressLogger(log.Log).SetEvent("block")

	cfg := BlockWatcherConfig{
		MaxQueue:     config.GetInt("max_queue"),
		DialTimeout:  config.GetDuration("zmq.dial_timeout"),
		RetryTimeout: config.GetDuration("zmq.retry_timeout"),
		ReadTimeout:  config.GetDuration("zmq.read_timeout"),
	}

	for {
		// check for shutdown
		select {
		case <-sctx.Done():
			return nil
		default:
		}

		// catch-up providers
		for _, p := range provider.Providers() {
			if !p.Enabled() {
				continue
			}
			log.Infof("Filling history for %s provider from block %d...", p.Name(), p.Status().Height)
			if err := p.FillHistory(ctx); err != nil {
				log.Error(err)
			}
		}

		// (re)connect
		if bw == nil {
			log.Info("Connecting block watcher...")
			cfg.StartBlock = provider.MinStatus()
			log.Infof("Crawling from block %d %s %s...",
				cfg.StartBlock.Height,
				cfg.StartBlock.Hash,
				cfg.StartBlock.Time,
			)
			bw = NewBlockWatcher(cfg, c)
			if err := bw.Connect(sctx, zmqurl); err != nil {
				log.Error(err)
				bw.Close()
				bw = nil
				// try reconnect
				select {
				case <-sctx.Done():
					return nil
				case <-time.After(cfg.RetryTimeout):
					continue
				}
			}
		}

		// process blocks
	process:
		for {
			select {
			case <-sctx.Done():
				return sctx.Err()
			case block := <-bw.Chan():
				if block == nil {
					// reconnect
					log.Info("Reconnect on close...")
					bw.Close()
					bw = nil
					break process
				}
				plog.Log(1, fmt.Sprintf("%d %s", block.Height, block.Timestamp))
				// skip blocks without relevant transactions
				if block.NContractCalls+block.NewContracts == 0 {
					continue
				}
				log.Debugf("Processing block %d %s...", block.Height, block.Hash)
				ops, err := fetchAllBlockOps(ctx, c, block.Hash)
				if err == nil {
					block.Ops = ops
					// forward to providers
					for _, p := range provider.Providers() {
						if !p.Enabled() {
							continue
						}
						// reorg-safe with tendermint
						// err = p.DisconnectBlock(sctx, block)
						err = p.ConnectBlock(sctx, block)
						if err != nil {
							break
						}
					}
				}
				if err != nil {
					log.Error(err)
					// reconnect
					bw.Close()
					bw = nil
					break process
				}
			}
		}

		// stop on shutdown
		select {
		case <-sctx.Done():
			return nil
		default:
		}
	}
}
