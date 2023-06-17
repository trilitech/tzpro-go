package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/tzpro"
	"github.com/echa/log"
)

var (
	api     string
	verbose bool
	vdebug  bool
	vtrace  bool
	mode    string
)

func init() {
	flag.StringVar(&api, "api", "https://api.tzpro.io", "use API")
	flag.BoolVar(&verbose, "v", false, "verbose")
	flag.BoolVar(&vdebug, "vv", false, "debug")
	flag.BoolVar(&vtrace, "vvv", false, "trace")
	flag.StringVar(&mode, "mode", "explorer", "use `explorer` or `table` mode")
}

func main() {
	flag.Parse()
	switch {
	case verbose:
		log.SetLevel(log.LevelInfo)
	case vdebug:
		log.SetLevel(log.LevelDebug)
	case vtrace:
		log.SetLevel(log.LevelTrace)
	}
	// create a new SDK client
	c, err := tzpro.NewClient(api, nil)
	if err != nil {
		log.Fatal(err)
	}
	c.WithLogger(log.Log)

	switch mode {
	case "explorer":
		err = runExplorer(c)
	case "table":
		err = runTable(c)
	default:
		log.Fatal("Invalid mode", mode)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func runExplorer(c *tzpro.Client) error {
	// use explorer API to get account info and embed metadata if available
	a, err := c.GetAccount(
		context.Background(),
		tezos.MustParseAddress(os.Args[1]),
		tzpro.NewAccountParams().WithMeta(),
	)
	if err != nil {
		return err
	}
	buf, _ := json.MarshalIndent(a, "", "  ")
	fmt.Println(string(buf))
	return nil
}

func runTable(c *tzpro.Client) error {
	// use table API to get raw account info
	q := c.NewAccountQuery()
	q.WithFilter(tzpro.FilterModeEqual, "address", os.Args[1])
	res, err := q.Run(context.Background())
	if err != nil {
		return err
	}
	a := res.Rows[0]
	buf, _ := json.MarshalIndent(a, "", "  ")
	fmt.Println(string(buf))
	return nil
}
