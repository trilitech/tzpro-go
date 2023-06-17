package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"blockwatch.cc/tzpro-go/tzpro"
	"github.com/echa/log"
)

var (
	api     string
	verbose bool
	vdebug  bool
	vtrace  bool
)

func init() {
	flag.StringVar(&api, "api", "https://api.tzpro.io", "use API")
	flag.BoolVar(&verbose, "v", false, "verbose")
	flag.BoolVar(&vdebug, "vv", false, "debug")
	flag.BoolVar(&vtrace, "vvv", false, "trace")
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
	if err := run(); err != nil {
		fmt.Println("Error:", err)
	}
}

func run() error {
	// use a placeholder calling context
	ctx := context.Background()

	// create a new SDK client
	c, err := tzpro.NewClient(api, nil)
	if err != nil {
		return err
	}

	// fetch block
	q := c.NewBlockQuery()
	q.WithLimit(1).WithDesc()
	res, err := q.Run(ctx)
	if err != nil {
		return err
	}

	buf, _ := json.MarshalIndent(res.Rows[0], "", "  ")
	fmt.Println(string(buf))
	return nil
}
