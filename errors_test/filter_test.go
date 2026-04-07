package errors_test

import (
	"errors"
	"testing"

	"github.com/ver1619/itFrame/core"
	iterr "github.com/ver1619/itFrame/errors"
)

func TestFilter_ErrorPassThrough(t *testing.T) {
	it := core.Slice([]iterr.Result[int]{
		iterr.ErrResult[int](errors.New("fail")),
	})

	f := iterr.Filter(it, func(x int) bool { return x > 0 })

	r, _ := f.Next()

	if r.Err == nil {
		t.Fatal("error should pass through filter")
	}
}
