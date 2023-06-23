// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"blockwatch.cc/tzgo/micheline"
	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go/tzpro"
	ct "github.com/daviddengcn/go-colortext"
	"github.com/echa/log"
)

var (
	flags      = flag.NewFlagSet("contract", flag.ContinueOnError)
	withPrim   bool
	withUnpack bool
	nocolor    bool
	node       string
	api        string
	verbose    bool
	vdebug     bool
	vtrace     bool
)

func init() {
	flags.Usage = func() {}
	flags.BoolVar(&verbose, "v", false, "verbose")
	flags.BoolVar(&vdebug, "vv", false, "debug")
	flags.BoolVar(&vtrace, "vvv", false, "trace")
	flags.StringVar(&node, "node", "https://rpc.tzpro.io", "Tezos node url")
	flags.StringVar(&api, "api", "https://api.tzpro.io", "TzPro API url")
	flags.BoolVar(&withPrim, "prim", false, "show primitives")
	flags.BoolVar(&withUnpack, "unpack", false, "unpack packed contract data")
	flags.BoolVar(&nocolor, "no-color", false, "disable color output")
}

func main() {
	if err := flags.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			fmt.Printf("Usage: contract [options] <cmd> <contract|operation>\n\n")
			flags.PrintDefaults()
			fmt.Println("\nCommands:")
			fmt.Println("  info      Show `contract` info")
			fmt.Println("  type      Show `contract` type")
			fmt.Println("  entry     Show `contract` entrypoints")
			fmt.Println("  storage   Show `contract` storage")
			fmt.Println("  params    Show call parameters from `operation`")
			os.Exit(0)
		}
		log.Fatal("Error:", err)
		os.Exit(1)
	}
	switch {
	case verbose:
		log.SetLevel(log.LevelInfo)
	case vdebug:
		log.SetLevel(log.LevelDebug)
	case vtrace:
		log.SetLevel(log.LevelTrace)
	}
	if err := run(); err != nil {
		if e, ok := tzpro.IsErrApi(err); ok {
			fmt.Printf("Error: %s: %s\n", e.Errors[0].Message, e.Errors[0].Detail)
		} else {
			fmt.Printf("Error: %v\n", err)
		}
		os.Exit(1)
	}
}

func run() error {
	if flags.NArg() < 1 {
		return fmt.Errorf("command required")
	}
	if flags.NArg() < 2 {
		return fmt.Errorf("argument required")
	}
	cmd := flags.Arg(0)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := tzpro.NewClient(api, nil).WithLogger(log.Log)

	switch cmd {
	case "info":
		addr, err := tezos.ParseAddress(flags.Arg(1))
		if err != nil {
			return err
		}
		return getContractInfo(ctx, c, addr)
	case "type":
		addr, err := tezos.ParseAddress(flags.Arg(1))
		if err != nil {
			return err
		}
		return getContractType(ctx, c, addr)
	case "entry":
		addr, err := tezos.ParseAddress(flags.Arg(1))
		if err != nil {
			return err
		}
		return getContractEntrypoints(ctx, c, addr)
	case "storage":
		addr, err := tezos.ParseAddress(flags.Arg(1))
		if err != nil {
			return err
		}
		return getContractStorage(ctx, c, addr)
	case "params":
		oh, err := tezos.ParseOpHash(flags.Arg(1))
		if err != nil {
			return err
		}
		return getContractCall(ctx, c, oh)
	default:
		return fmt.Errorf("unsupported command %s", cmd)
	}
}

func getContractInfo(ctx context.Context, c *tzpro.Client, addr tezos.Address) error {
	cc, err := c.Contract.Get(ctx, addr, tzpro.NewParams())
	if err != nil {
		return err
	}
	fmt.Println("Contract Info:")
	print(cc, 2)
	return nil
}

func getContractType(ctx context.Context, c *tzpro.Client, addr tezos.Address) error {
	script, err := c.Contract.GetScript(ctx, addr, tzpro.NewParams().WithPrim())
	if err != nil {
		return err
	}
	fmt.Println("Storage Type:")
	print(script.Script.StorageType(), 2)
	if withPrim {
		fmt.Println("Michelson:")
		print(script.Script.StorageType().Prim, 0)
	}
	return nil
}

func getContractEntrypoints(ctx context.Context, c *tzpro.Client, addr tezos.Address) error {
	cc, err := c.Contract.GetScript(ctx, addr, tzpro.NewParams().WithPrim())
	if err != nil {
		return err
	}
	fmt.Println("Entrypoints:")
	eps, err := cc.Script.Entrypoints(withPrim)
	if err != nil {
		return err
	}
	print(eps, 2)
	// if withPrim {
	// 	fmt.Println("Michelson:")
	// 	print(cc.Script.StorageType().Prim, 0)
	// }
	return nil
}

func getContractStorage(ctx context.Context, c *tzpro.Client, addr tezos.Address) error {
	p := tzpro.NewParams().WithPrim()
	cc, err := c.Contract.GetScript(ctx, addr, p)
	if err != nil {
		return err
	}
	store, err := c.Contract.GetStorage(ctx, addr, p)
	if err != nil {
		return err
	}
	fmt.Println("Storage Contents:")
	print(micheline.NewValue(cc.Script.StorageType(), *store.Prim), 2)
	if withPrim {
		fmt.Println("Michelson:")
		print(store.Prim, 0)
	}
	return nil
}

func getContractCall(ctx context.Context, c *tzpro.Client, hash tezos.OpHash) error {
	ops, err := c.Op.Get(ctx, hash, tzpro.NewParams().WithPrim())
	if err != nil {
		return err
	}
	for i, op := range ops {
		fmt.Printf("%s %d %d/%d\n", op.Hash, i, op.OpN, op.OpP)
		fmt.Printf("  type:    %s\n", op.Type)
		fmt.Printf("  target:  %s\n", op.Receiver)
		fmt.Printf("  params:  %t\n", op.HasParameters())
		fmt.Printf("  storage: %t\n", op.HasStorage())
		fmt.Printf("  bigmap:  %t\n", op.HasBigmapUpdates())
		if op.Type != tzpro.OpTypeTransaction {
			continue
		}
		if !op.IsContract {
			continue
		}
		if op.HasParameters() {
			script, err := c.Contract.GetScript(ctx, op.Receiver, tzpro.NewParams().WithPrim())
			if err != nil {
				return err
			}
			op.WithScript(script)
			args, err := op.DecodeParams(false, 0)
			if err != nil {
				return err
			}
			fmt.Printf("  Entrypoint: %s\n", args.Entrypoint)
			fmt.Printf("  Args: %#v", args)
			fmt.Println("  Params:")
			print(args.Value, 2)
			if withPrim {
				fmt.Println("Michelson:")
				print(args.Prim, 0)
			}
		} else {
			fmt.Println("  Transfer only")
		}
	}
	return nil
}

// Color print helpers
func print(val interface{}, n int) error {
	var (
		body []byte
		err  error
	)
	if n > 0 {
		body, err = json.MarshalIndent(val, "", strings.Repeat(" ", n))
	} else {
		body, err = json.Marshal(val)
	}
	if err != nil {
		return err
	}
	if nocolor {
		os.Stdout.Write(body)
	} else {
		var raw interface{}
		// raw := make(map[string]interface{})
		dec := json.NewDecoder(bytes.NewBuffer(body))
		dec.UseNumber()
		dec.Decode(&raw)
		printJSON(1, raw, false)
	}
	fmt.Println()
	return nil
}

func printJSON(depth int, val interface{}, isKey bool) {
	switch v := val.(type) {
	case nil:
		ct.ChangeColor(ct.Blue, false, ct.None, false)
		fmt.Print("null")
		ct.ResetColor()
	case bool:
		ct.ChangeColor(ct.Blue, false, ct.None, false)
		if v {
			fmt.Print("true")
		} else {
			fmt.Print("false")
		}
		ct.ResetColor()
	case string:
		if isKey {
			ct.ChangeColor(ct.Blue, true, ct.None, false)
		} else {
			ct.ChangeColor(ct.Yellow, false, ct.None, false)
		}
		fmt.Print(strconv.Quote(v))
		ct.ResetColor()
	case json.Number:
		ct.ChangeColor(ct.Blue, false, ct.None, false)
		fmt.Print(v)
		ct.ResetColor()
	case map[string]interface{}:

		if len(v) == 0 {
			fmt.Print("{}")
			break
		}

		var keys []string

		for h := range v {
			keys = append(keys, h)
		}

		sort.Strings(keys)

		fmt.Println("{")
		needNL := false
		for _, key := range keys {
			if needNL {
				fmt.Print(",\n")
			}
			needNL = true
			for i := 0; i < depth; i++ {
				fmt.Print("    ")
			}

			printJSON(depth+1, key, true)
			fmt.Print(": ")
			printJSON(depth+1, v[key], false)
		}
		fmt.Println("")

		for i := 0; i < depth-1; i++ {
			fmt.Print("    ")
		}
		fmt.Print("}")

	case []interface{}:

		if len(v) == 0 {
			fmt.Print("[]")
			break
		}

		fmt.Println("[")
		needNL := false
		for _, e := range v {
			if needNL {
				fmt.Print(",\n")
			}
			needNL = true
			for i := 0; i < depth; i++ {
				fmt.Print("    ")
			}

			printJSON(depth+1, e, false)
		}
		fmt.Println("")

		for i := 0; i < depth-1; i++ {
			fmt.Print("    ")
		}
		fmt.Print("]")
	default:
		fmt.Println("unknown type:", reflect.TypeOf(v))
	}
}
