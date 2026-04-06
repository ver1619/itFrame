package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	it := ops.CrossJoin(
		core.Slice([]int{1, 2}),
		core.Slice([]string{"a", "b"}),
	)

	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
