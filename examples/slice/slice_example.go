package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
)

func main() {
	it := core.NewSliceIterator([]int{10, 20, 30})

	for {
		val, ok := it.Next()
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
