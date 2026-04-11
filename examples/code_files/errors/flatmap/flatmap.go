package main

import (
	"errors"
	"fmt"

	"github.com/ver1619/itFrame/core"
	iterr "github.com/ver1619/itFrame/errors"
)

func main() {
	stream := iterr.FromSlice([]iterr.Result[int]{
		iterr.Ok(2),
		iterr.ErrResult[int](errors.New("bad input")),
	})

	result, err := stream.FlatMap(
		func(x int) core.Iterator[iterr.Result[int]] {
			return core.Slice([]iterr.Result[int]{
				iterr.Ok(x),
				iterr.Ok(x * x),
			})
		},
	).Collect()

	fmt.Println(result)
	fmt.Println(err)
}
