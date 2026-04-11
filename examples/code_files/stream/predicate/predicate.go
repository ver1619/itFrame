package main

import (
	"fmt"

	"github.com/ver1619/itFrame/stream"
)

func main() {
	ok := stream.Range(0, 10, 1).
		Any(func(x int) bool { return x == 7 })

	fmt.Println(ok)
}
