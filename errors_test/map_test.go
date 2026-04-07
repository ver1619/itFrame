package errors_test

import (
	"errors"
	"testing"

	"github.com/ver1619/itFrame/core"
	iterr "github.com/ver1619/itFrame/errors"
)

func TestMap_ErrorPropagation(t *testing.T) {
	it := core.Slice([]iterr.Result[int]{
		iterr.Ok(1),
		iterr.ErrResult[int](errors.New("fail")),
	})

	mapped := iterr.Map(it, func(x int) int { return x * 2 })

	r1, _ := mapped.Next()
	r2, _ := mapped.Next()

	if r1.Value != 2 {
		t.Fatal("expected mapped value")
	}

	if r2.Err == nil {
		t.Fatal("error should propagate")
	}
}
