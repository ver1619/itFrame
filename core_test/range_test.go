package core_test

import (
	"testing"

	"github.com/ver1619/itFrame/core"
)

func TestRangeIterator_Forward(t *testing.T) {
	it := core.NewRangeIterator(0, 5, 1)

	expected := []int{0, 1, 2, 3, 4}
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

func TestRangeIterator_Backward(t *testing.T) {
	it := core.NewRangeIterator(5, 0, -1)

	expected := []int{5, 4, 3, 2, 1}
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

func TestRangeIterator_InvalidForward(t *testing.T) {
	it := core.NewRangeIterator(5, 0, 1)

	_, ok := it.Next()
	if ok {
		t.Fatal("expected empty iterator for invalid forward range")
	}
}

func TestRangeIterator_InvalidBackward(t *testing.T) {
	it := core.NewRangeIterator(0, 5, -1)

	_, ok := it.Next()
	if ok {
		t.Fatal("expected empty iterator for invalid backward range")
	}
}

func TestRangeIterator_ExhaustionStability(t *testing.T) {
	it := core.NewRangeIterator(0, 1, 1)

	_, _ = it.Next()

	for i := 0; i < 5; i++ {
		_, ok := it.Next()
		if ok {
			t.Fatal("expected stable exhaustion behavior")
		}
	}
}

func TestRangeIterator_StepZeroPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for step = 0")
		}
	}()

	core.NewRangeIterator(0, 10, 0)
}
