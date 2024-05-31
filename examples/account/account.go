package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/echa/log"
	"github.com/trilitech/tzpro-go/tzpro"
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
	c := tzpro.NewClient(api, nil).WithLogger(log.Log)

	var err error
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
	a, err := c.Account.Get(
		context.Background(),
		tzpro.NewAddress(flag.Arg(0)),
		tzpro.WithMeta(),
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
	q := c.Account.NewQuery().AndEqual("address", flag.Arg(0))
	res, err := q.Run(context.Background())
	if err != nil {
		return err
	}
	buf, _ := json.MarshalIndent(res.Rows()[0], "", "  ")
	fmt.Println(res.Cursor(), string(buf))
	return nil
}
