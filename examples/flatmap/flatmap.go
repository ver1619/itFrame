package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	it := ops.FlatMap(
		core.Slice([]int{1, 2, 3}),
		func(x int) core.Iterator[int] {
			return core.Slice([]int{x, x * x})
		},
	)

	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		fmt.Print(v, " ")
	}
}
