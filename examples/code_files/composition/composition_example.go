package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	it := ops.Filter(
		ops.Map(
			core.Slice([]int{1, 2, 3, 4}),
			func(x int) int { return x * x },
		),
		func(x int) bool { return x%2 == 0 },
	)

	for {
		val, ok := it.Next()
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
