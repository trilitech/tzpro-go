package main

import (
	"context"
	"fmt"

	"blockwatch.cc/tzpro-go/tzpro"

	"github.com/echa/log"
)

func main() {
	log.SetLevel(log.LevelDebug)
	c, _ := tzpro.NewClient("http://localhost:8000", nil)
	c.WithLogger(log.Log)
	q := c.NewContractQuery()
	_, err := q.Run(context.Background())
	fmt.Println(err)
}
