package core_test

import (
	"testing"

	"github.com/ver1619/itFrame/core"
)

func TestSliceIterator_Normal(t *testing.T) {
	it := core.Slice([]int{1, 2, 3})

	expected := []int{1, 2, 3}
	for i, v := range expected {
		val, ok := it.Next()
		if !ok {
			t.Fatalf("expected value at index %d", i)
		}
		if val != v {
			t.Fatalf("expected %d, got %d", v, val)
		}
	}

	_, ok := it.Next()
	if ok {
		t.Fatal("expected iterator to be exhausted")
	}
}

func TestSliceIterator_Empty(t *testing.T) {
	it := core.Slice([]int{})

	_, ok := it.Next()
	if ok {
		t.Fatal("expected empty iterator")
	}
}

func TestSliceIterator_Nil(t *testing.T) {
	var data []int
	it := core.Slice(data)

	_, ok := it.Next()
	if ok {
		t.Fatal("expected nil slice to behave as empty")
	}
}

func TestSliceIterator_ExhaustionStability(t *testing.T) {
	it := core.Slice([]int{1})

	_, _ = it.Next()

	for i := 0; i < 5; i++ {
		_, ok := it.Next()
		if ok {
			t.Fatal("expected stable exhaustion behavior")
		}
	}
}
