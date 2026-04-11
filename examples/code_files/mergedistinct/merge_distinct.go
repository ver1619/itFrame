package main

import (
	"fmt"

	"github.com/ver1619/itFrame/compare"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	cmp := compare.LessFunc[int](func(a, b int) bool {
		return a < b
	})

	it := ops.MergeDistinct(
		core.Slice([]int{1, 2, 3, 5}),
		core.Slice([]int{2, 3, 4}),
		cmp,
	)

	fmt.Println(ops.Collect(it))
}
