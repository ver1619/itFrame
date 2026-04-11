package main

import (
	"fmt"

	"github.com/ver1619/itFrame/stream"
)

func main() {
	sum := stream.Range(0, 5, 1).
		Reduce(0, func(acc, x int) int { return acc + x })

	fmt.Println(sum)
}
