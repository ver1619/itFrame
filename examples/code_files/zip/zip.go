package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	it := ops.Zip(
		core.Slice([]int{1, 2, 3, 4}),
		core.Slice([]string{"a", "b", "c"}),
	)

	fmt.Println(ops.Collect(it))
}
