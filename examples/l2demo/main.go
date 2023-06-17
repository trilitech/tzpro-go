// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/tzpro"
	"github.com/echa/config"
	"github.com/echa/log"
)

var (
	flags    = flag.NewFlagSet("l2demo", flag.ContinueOnError)
	verbose  bool
	sorted   bool
	nobackup bool
	nocolor  bool
	apiurl   string
	zmqurl   string
)

func init() {
	flags.Usage = func() {}
	flags.BoolVar(&verbose, "v", false, "be verbose")
	flags.BoolVar(&nocolor, "no-color", false, "disable color output")
	flags.StringVar(&apiurl, "index", "http://127.0.0.1:8000", "Index API URL")
	flags.StringVar(&zmqurl, "zmq", "tcp://127.0.0.1:27000", "ZMQ publisher address")
}

func main() {
	if err := flags.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			fmt.Println("Usage: l2demo [address[/asset_id]]")
			flags.PrintDefaults()
			os.Exit(0)
		}
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	config.SetEnvPrefix("L2DEMO")
	realconf := config.ConfigName()
	if _, err := os.Stat(realconf); err == nil {
		if err := config.ReadConfigFile(); err != nil {
			fmt.Printf("Could not read config %s: %v\n", realconf, err)
			os.Exit(1)
		}
		log.Infof("Using configuration file %s", realconf)
	} else {
		log.Warnf("Missing config file, using default values.")
	}

	if verbose {
		log.SetLevel(log.LevelDebug)
	}

	if err := run(); err != nil {
		if e, ok := tzpro.IsApiError(err); ok {
			log.Errorf("%s: %s", e.Errors[0].Message, e.Errors[0].Detail)
		} else {
			log.Error(err)
		}
		os.Exit(1)
	}
}

func run() error {
	if flags.NArg() < 1 {
		return fmt.Errorf("contract address required")
	}
	addr, err := tezos.ParseAddress(flags.Arg(0))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := tzpro.NewClient(apiurl, nil)
	if err != nil {
		return err
	}
	client.WithLogger(log.Log)

	return watchContract(ctx, client, addr)
}
