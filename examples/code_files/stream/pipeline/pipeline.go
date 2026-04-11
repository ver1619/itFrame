package main

import (
	"fmt"

	"github.com/ver1619/itFrame/stream"
)

func main() {
	result := stream.Slice([]int{1, 2, 3, 4}).
		Map(func(x int) int { return x * x }).
		Filter(func(x int) bool { return x%2 == 0 }).
		Collect()

	fmt.Println(result)
}
