package ops_test

import (
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestReduce_Sum(t *testing.T) {
	it := core.Range(0, 5, 1)

	sum := ops.Reduce(it, 0, func(acc, x int) int {
		return acc + x
	})

	if sum != 10 {
		t.Fatalf("expected 10, got %d", sum)
	}
}

func TestReduce_Empty(t *testing.T) {
	it := core.Slice([]int{})

	result := ops.Reduce(it, 100, func(acc, x int) int {
		return acc + x
	})

	if result != 100 {
		t.Fatalf("expected init value, got %d", result)
	}
}
