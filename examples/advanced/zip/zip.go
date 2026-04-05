package main

import (
	"fmt"

	"github.com/ver1619/itFrame/advanced"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	it := advanced.Zip(
		core.Slice([]int{1, 2, 3}),
		core.Slice([]string{"a", "b", "c"}),
	)

	fmt.Println(ops.Collect(it))
}
