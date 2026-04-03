package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	c := ops.Count(
		core.Range(0, 10, 1),
	)

	fmt.Println(c)
}
