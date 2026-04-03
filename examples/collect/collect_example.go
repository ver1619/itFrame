package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	data := ops.Collect(
		ops.Filter(
			core.Range(0, 10, 1),
			func(x int) bool { return x%2 == 0 },
		),
	)

	fmt.Println(data)
}
