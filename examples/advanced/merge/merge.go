package main

import (
	"fmt"

	"github.com/ver1619/itFrame/advanced"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	it := advanced.Merge(
		core.Slice([]int{1, 3, 5}),
		core.Slice([]int{2, 4, 6}),
		func(a, b int) bool { return a < b },
	)

	fmt.Println(ops.Collect(it))
}
