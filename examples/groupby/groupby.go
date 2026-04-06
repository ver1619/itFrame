package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	it := ops.GroupBy(
		core.Slice([]int{1, 2, 3, 4}),
		func(x int) int { return x % 2 },
	)

	for {
		group, ok := it.Next()
		if !ok {
			break
		}
		fmt.Println(group.Key, group.Items)
	}
}
