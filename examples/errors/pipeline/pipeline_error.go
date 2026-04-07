package main

import (
	"errors"
	"fmt"

	iterr "github.com/ver1619/itFrame/errors"
)

func main() {
	stream := iterr.FromSlice([]iterr.Result[int]{
		iterr.Ok(1),
		iterr.ErrResult[int](errors.New("fail")),
	})

	result, err := stream.
		Map(func(x int) int { return x * 2 }).
		Collect()

	fmt.Println(result)
	fmt.Println(err)
}
