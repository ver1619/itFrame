package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

type Result struct {
	Key   int
	Count int
}

func main() {
	grouped := ops.GroupBy(
		core.Slice([]int{1, 2, 3, 4, 5}),
		func(x int) int { return x % 2 },
	)

	result := ops.Aggregate(
		grouped,
		func(g ops.Group[int, int]) Result {
			return Result{
				Key:   g.Key,
				Count: len(g.Items),
			}
		},
	)

	for {
		v, ok := result.Next()
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
