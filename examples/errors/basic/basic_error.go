package main

import (
	"errors"
	"fmt"

	"github.com/ver1619/itFrame/core"
	iterr "github.com/ver1619/itFrame/errors"
)

func main() {
	it := core.Slice([]iterr.Result[int]{
		iterr.Ok(1),
		iterr.ErrResult[int](errors.New("fail")),
	})

	result, err := iterr.Collect(it)

	fmt.Println(result) // nil
	fmt.Println(err)    // fail
}
