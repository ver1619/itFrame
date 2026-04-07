package errors_test

import (
	"errors"
	"testing"

	iterr "github.com/ver1619/itFrame/errors"
)

func TestStream_Map_Filter(t *testing.T) {
	stream := iterr.FromSlice([]iterr.Result[int]{
		iterr.Ok(1),
		iterr.Ok(2),
	})

	result, err := stream.
		Map(func(x int) int { return x * 2 }).
		Filter(func(x int) bool { return x > 2 }).
		Collect()

	if err != nil {
		t.Fatal(err)
	}

	if len(result) != 1 || result[0] != 4 {
		t.Fatal("unexpected result")
	}
}

func TestStream_ErrorStops(t *testing.T) {
	stream := iterr.FromSlice([]iterr.Result[int]{
		iterr.Ok(1),
		iterr.ErrResult[int](errors.New("fail")),
	})

	_, err := stream.Collect()

	if err == nil {
		t.Fatal("expected error")
	}
}
