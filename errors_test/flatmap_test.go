package errors_test

import (
	"errors"
	"testing"

	"github.com/ver1619/itFrame/core"
	iterr "github.com/ver1619/itFrame/errors"
)

func TestFlatMap_ErrorPropagation(t *testing.T) {
	it := core.Slice([]iterr.Result[int]{
		iterr.ErrResult[int](errors.New("fail")),
	})

	f := iterr.FromIterator(it).FlatMap(
		func(x int) core.Iterator[iterr.Result[int]] {
			return core.Slice([]iterr.Result[int]{iterr.Ok(x)})
		},
	)

	r, _ := f.Iterator().Next()

	if r.Err == nil {
		t.Fatal("error should propagate")
	}
}
