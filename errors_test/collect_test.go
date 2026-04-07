package errors_test

import (
	"errors"
	"testing"

	"github.com/ver1619/itFrame/core"
	iterr "github.com/ver1619/itFrame/errors"
)

func TestCollect_StopOnError(t *testing.T) {
	it := core.Slice([]iterr.Result[int]{
		iterr.Ok(1),
		iterr.ErrResult[int](errors.New("fail")),
		iterr.Ok(3),
	})

	result, err := iterr.Collect(it)

	if err == nil {
		t.Fatal("expected error")
	}

	if result != nil {
		t.Fatal("result should be nil on error")
	}
}
