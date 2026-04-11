package main

import (
	"fmt"

	"github.com/ver1619/itFrame/stream"
)

func main() {
	// MapTo: int → string (type-changing map)
	labels := stream.MapTo(
		stream.Slice([]int{1, 2, 3}),
		func(x int) string { return fmt.Sprintf("item_%d", x) },
	).Collect()

	fmt.Println("MapTo:", labels)

	// FlatMapTo: int → []string (type-changing flatmap)
	expanded := stream.FlatMapTo(
		stream.Slice([]int{1, 2}),
		func(x int) []string {
			return []string{
				fmt.Sprintf("%d-a", x),
				fmt.Sprintf("%d-b", x),
			}
		},
	).Collect()

	fmt.Println("FlatMapTo:", expanded)

	// Chaining: filter → MapTo
	result := stream.MapTo(
		stream.Slice([]int{1, 2, 3, 4, 5}).
			Filter(func(x int) bool { return x > 2 }),
		func(x int) string { return fmt.Sprintf("[%d]", x) },
	).Collect()

	fmt.Println("Chained:", result)
}
