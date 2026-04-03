package ops_test

import (
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestAny(t *testing.T) {
	it := core.Range(0, 10, 1)

	ok := ops.Any(it, func(x int) bool { return x == 5 })

	if !ok {
		t.Fatal("expected true")
	}
}

func TestAll(t *testing.T) {
	it := core.Slice([]int{2, 4, 6})

	ok := ops.All(it, func(x int) bool { return x%2 == 0 })

	if !ok {
		t.Fatal("expected true")
	}
}

func TestAny_Empty(t *testing.T) {
	it := core.Slice([]int{})

	ok := ops.Any(it, func(x int) bool { return true })

	if ok {
		t.Fatal("expected false")
	}
}

func TestAll_Empty(t *testing.T) {
	it := core.Slice([]int{})

	ok := ops.All(it, func(x int) bool { return false })

	if !ok {
		t.Fatal("expected true for empty iterator")
	}
}
