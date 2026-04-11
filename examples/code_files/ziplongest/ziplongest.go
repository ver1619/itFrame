package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	it := ops.ZipLongest(
		core.Slice([]int{1, 2, 3}),
		core.Slice([]string{"a", "b"}),
	)

	fmt.Println(ops.Collect(it))
}
