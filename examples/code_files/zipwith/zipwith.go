package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	it := ops.ZipWith(
		core.Slice([]int{1, 2, 3}),
		core.Slice([]int{4, 5, 6}),
		func(a, b int) int { return a + b },
	)

	fmt.Println(ops.Collect(it))
}
