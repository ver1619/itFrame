package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	sum := ops.Reduce(
		core.Range(0, 5, 1),
		0,
		func(acc, x int) int { return acc + x },
	)

	fmt.Println(sum)
}
