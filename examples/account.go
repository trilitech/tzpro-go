package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"blockwatch.cc/tzgo/tezos"
	"blockwatch.cc/tzpro-go"
)

func main() {
	err := runExplorer()
	// err := runTable()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func runExplorer() error {
	// use explorer API to get account info and embed metadata if available
	a, err := tzpro.DefaultClient.GetAccount(
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

func runTable() error {
	// use table API to get raw account info
	q := tzpro.DefaultClient.NewAccountQuery()
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
