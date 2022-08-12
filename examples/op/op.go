package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"blockwatch.cc/tzpro-go"
)

func main() {
	// use default Mainnet client
	c, _ := tzpro.NewClient("https://api.tzpro.io", nil)
	ctx := context.Background()

	q := c.NewOpQuery()
	q.WithFilter(tzpro.FilterModeEqual, "hash", os.Args[1])
	res, err := q.Run(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		buf, _ := json.MarshalIndent(res.Rows[0], "", "  ")
		fmt.Println(string(buf))
	}
}
