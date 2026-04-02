package ops_test

import (
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestFilterIterator_Basic(t *testing.T) {
	it := ops.Filter(
		core.Slice([]int{1, 2, 3, 4}),
		func(x int) bool { return x%2 == 0 },
	)

	expected := []int{2, 4}

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

func TestFilterIterator_AllFiltered(t *testing.T) {
	it := ops.Filter(
		core.Slice([]int{1, 3, 5}),
		func(x int) bool { return false },
	)

	_, ok := it.Next()
	if ok {
		t.Fatal("expected no elements")
	}
}

func TestFilterIterator_NoneFiltered(t *testing.T) {
	it := ops.Filter(
		core.Slice([]int{1, 2}),
		func(x int) bool { return true },
	)

	expected := []int{1, 2}

	for i, v := range expected {
		val, ok := it.Next()
		if !ok || val != v {
			t.Fatalf("expected %d at index %d", v, i)
		}
	}
}
