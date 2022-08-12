package main

import (
	"blockwatch.cc/tzpro-go"
	"context"
	"fmt"

	"github.com/echa/log"
)

func main() {
	log.SetLevel(log.LevelDebug)
	tzpro.UseLogger(log.Log)
	c, _ := tzpro.NewClient("http://localhost:8000", nil)
	q := c.NewContractQuery()
	_, err := q.Run(context.Background())
	fmt.Println(err)
}
