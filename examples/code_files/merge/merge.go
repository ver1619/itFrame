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

	it := ops.Merge(
		core.Slice([]int{1, 3, 5}),
		core.Slice([]int{2, 4, 6}),
		cmp,
	)

	result := ops.Collect(it)

	fmt.Println(result)
}
