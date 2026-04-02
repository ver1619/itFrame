package ops_test

import (
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestMapIterator_Basic(t *testing.T) {
	it := ops.Map(
		core.Slice([]int{1, 2, 3}),
		func(x int) int { return x * 2 },
	)

	expected := []int{2, 4, 6}

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

func TestMapIterator_Empty(t *testing.T) {
	it := ops.Map(
		core.Slice([]int{}),
		func(x int) int { return x * 2 },
	)

	_, ok := it.Next()
	if ok {
		t.Fatal("expected empty iterator")
	}
}

func TestMapIterator_Identity(t *testing.T) {
	it := ops.Map(
		core.Slice([]int{5, 6}),
		func(x int) int { return x },
	)

	expected := []int{5, 6}

	for i, v := range expected {
		val, ok := it.Next()
		if !ok || val != v {
			t.Fatalf("expected %d at index %d", v, i)
		}
	}
}
