package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	exists := ops.Any(
		core.Range(0, 10, 1),
		func(x int) bool { return x == 7 },
	)

	fmt.Println(exists)
}
