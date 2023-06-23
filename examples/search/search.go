// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

// Search contract calls for address used in parameters
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/tzpro"
	"github.com/echa/log"
)

var (
	flags   = flag.NewFlagSet("search", flag.ContinueOnError)
	verbose bool
	node    string
	index   string
)

func init() {
	flags.Usage = func() {}
	flags.BoolVar(&verbose, "v", false, "be verbose")
	flags.StringVar(&index, "index", "https://api.tzpro.io", "TzPro API Url")
}

func main() {
	if err := flags.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			fmt.Println("Tezos Contract Call Search")
			flags.PrintDefaults()
			os.Exit(0)
		}
		log.Fatal("Error:", err)
	}

	if verbose {
		log.SetLevel(log.LevelDebug)
	}

	if err := run(); err != nil {
		log.Fatal("Error:", err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := tzpro.NewClient(index, nil).WithLogger(log.Log)

	if err := searchCalls(ctx, c); err != nil {
		return err
	}

	return nil
}

// Using Explorer API
func searchCalls(ctx context.Context, c *tzpro.Client) error {
	recv, err := tezos.ParseAddress(flags.Arg(0))
	if err != nil {
		return err
	}
	addr, err := tezos.ParseAddress(flags.Arg(1))
	if err != nil {
		return err
	}
	log.Infof("Searching calls to %s for address %s", recv, addr)

	p := tzpro.NewParams().WithLimit(500)
	plog := log.NewProgressLogger(log.Log)
	var (
		count int
	)
	for {
		calls, err := c.Contract.ListCalls(ctx, recv, p)
		if err != nil {
			return err
		}
		if len(calls) == 0 {
			break
		}
		for _, v := range calls {
			found := false
			if v.Parameters != nil {
				args, _ := v.DecodeParams(false, 0)
				err := args.Walk("", func(path string, value interface{}) error {
					if value == nil {
						return nil
					}
					if s, ok := value.(tezos.Address); ok {
						found = found || s.Equal(addr)
					}
					return nil
				})
				if err != nil {
					log.Errorf("%s: %v", v.Hash, err)
				}
			}
			if v.Storage != nil {
				store, _ := v.DecodeStorage(false, 0)
				err := store.Walk("", func(path string, value interface{}) error {
					if value == nil {
						return nil
					}
					if s, ok := value.(tezos.Address); ok {
						found = found || s.Equal(addr)
					}
					return nil
				})
				if err != nil {
					log.Errorf("%s: %v", v.Hash, err)
				}
			}
			events, _ := v.DecodeBigmapUpdates(false, false, 0)
			for _, bmd := range events {
				err := bmd.Walk("", func(path string, value interface{}) error {
					if value == nil {
						return nil
					}
					if s, ok := value.(tezos.Address); ok {
						found = found || s.Equal(addr)
					}
					return nil
				})
				if err != nil {
					log.Errorf("%s: %v", v.Hash, err)
				}
			}
			count++
			if found {
				log.Infof("%s matches", v.Hash)
			}
		}
		plog.Log(len(calls))
		p = p.WithCursor(calls[len(calls)-1].Id)
	}
	log.Infof("Processed %d calls", count)
	return nil
}

// Using Table API
func search(ctx context.Context, c *tzpro.Client) error {
	recv := flags.Arg(0)
	addr := flags.Arg(1)
	log.Infof("Searching calls to %s for address %s", recv, addr)

	q := c.Op.NewQuery().
		WithLimit(50000).
		WithEqual("type", "transaction").
		WithEqual("receiver", recv).
		WithColumns("id", "hash", "parameters", "big_map_diff", "is_contract", "receiver")

	plog := log.NewProgressLogger(log.Log)
	var (
		cursor uint64
		count  int
	)
	for {
		q.WithCursor(cursor)
		log.Infof("Fetching calls from id %d...", cursor)
		ops, err := q.Run(ctx)
		if err != nil {
			return err
		}
		if ops.Len() == 0 {
			break
		}
		if err := c.Op.ResolveTypes(ctx, ops.Rows()...); err != nil {
			return err
		}
		for _, v := range ops.Rows() {
			found := false
			if v.Parameters != nil {
				args, err := v.DecodeParams(false, 0)
				if err != nil {
					log.Errorf("%s: %v", v.Hash, err)
					continue
				}
				err = args.Walk("", func(path string, value interface{}) error {
					if value == nil {
						return nil
					}
					if s, ok := value.(string); ok {
						found = found || s == addr
					}
					log.Infof("%s: param %s = %v", v.Hash, path, value)
					return nil
				})
				if err != nil {
					log.Errorf("%s: %v", v.Hash, err)
				}
			}
			// deprecated on API
			// if v.Storage != nil {
			// 	store, _ := v.DecodeStorage(false, 0)
			// 	err := store.Walk("", func(path string, value interface{}) error {
			// 		if value == nil {
			// 			return nil
			// 		}
			// 		log.Infof("%s: storage %s = %v", v.Hash, path, value)
			// 		if s, ok := value.(string); ok {
			// 			found = found || s == addr
			// 		}
			// 		return nil
			// 	})
			// 	if err != nil {
			// 		log.Errorf("%s: %v", v.Hash, err)
			// 	}
			// }
			events, _ := v.DecodeBigmapUpdates(false, false, 0)
			for _, bmd := range events {
				err := bmd.Walk("", func(path string, value interface{}) error {
					if value == nil {
						return nil
					}
					log.Infof("%s: bigmap %s = %v", v.Hash, path, value)
					if s, ok := value.(string); ok {
						found = found || s == addr
					}
					return nil
				})
				if err != nil {
					log.Errorf("%s: %v", v.Hash, err)
				}
			}
			count++
			if found {
				log.Infof("%s matches", v.Hash)
			}
		}
		plog.Log(ops.Len())
		cursor = ops.Cursor()
	}
	log.Infof("Processed %d calls", count)
	return nil
}
