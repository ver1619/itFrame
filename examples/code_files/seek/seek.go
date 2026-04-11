package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	it := ops.Seek(
		core.Range(0, 10, 1),
		func(x int) bool { return x >= 5 },
	)

	fmt.Println(ops.Collect(it))
}
