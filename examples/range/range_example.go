package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
)

func main() {
	it := core.Range(0, 5, 1)

	for {
		val, ok := it.Next()
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
