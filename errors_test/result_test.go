package errors_test

import (
	"errors"
	"testing"

	iterr "github.com/ver1619/itFrame/errors"
)

func TestResult_Ok(t *testing.T) {
	r := iterr.Ok(10)

	if r.Err != nil || r.Value != 10 {
		t.Fatal("expected Ok result")
	}
}

func TestResult_Error(t *testing.T) {
	err := errors.New("fail")
	r := iterr.ErrResult[int](err)

	if r.Err == nil {
		t.Fatal("expected error")
	}
}
