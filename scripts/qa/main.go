package main

import (
	"context"
	"flag"
	"fmt"
	// "runtime/debug"
	"strings"

	"blockwatch.cc/tzpro-go/tzpro"
	"github.com/echa/log"
)

var (
	nFail   int
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

func try(name string, fn func()) {
	fmt.Printf("%s %s ", name, strings.Repeat(".", 26-len(name)))
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("FAILED\nError: %v\n", err)
			nFail++
			// fmt.Println(string(debug.Stack()))
		} else {
			fmt.Println("OK")
		}
	}()
	fn()
}

func run() error {
	// use a placeholder calling context
	ctx := context.Background()

	// create a new SDK client
	c := tzpro.NewClient(api, nil).WithLogger(log.Log)

	tip := TestCommon(ctx, c)
	TestBlock(ctx, c, tip)
	TestWallet(ctx, c)
	TestBaker(ctx, c)
	TestContract(ctx, c)
	TestMarket(ctx, c)
	TestToken(ctx, c)
	TestDex(ctx, c)
	TestFarm(ctx, c)
	TestLend(ctx, c)
	TestNft(ctx, c)

	if nFail > 0 {
		fmt.Printf("%d tests have FAILED.", nFail)
	} else {
		fmt.Println("All tests have PASSED.")
	}
	return nil
}
