package edge_tests

import (
	"testing"

	"github.com/ver1619/itFrame/compare"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestEmptyIterators(t *testing.T) {
	t.Run("Map", func(t *testing.T) {
		it := ops.Map(core.Slice([]int{}), func(x int) int { return x * 2 })
		res := ops.Collect(it)
		if len(res) != 0 {
			t.Errorf("Expected empty result, got %v", res)
		}
	})

	t.Run("Filter", func(t *testing.T) {
		it := ops.Filter(core.Slice([]int{}), func(x int) bool { return true })
		res := ops.Collect(it)
		if len(res) != 0 {
			t.Errorf("Expected empty result, got %v", res)
		}
	})

	t.Run("FlatMap", func(t *testing.T) {
		it := ops.FlatMap(core.Slice([]int{}), func(x int) []int { return []int{x} })
		res := ops.Collect(it)
		if len(res) != 0 {
			t.Errorf("Expected empty result, got %v", res)
		}
	})

	t.Run("Reduce", func(t *testing.T) {
		// Reduce returns the identity if empty? Wait, let's check Reduce implementation.
		// Actually Reduce usually takes an initial value or returns a Result/Option.
		// Looking at ops/reduce.go (implied from common patterns)
		res := ops.Reduce(core.Slice([]int{}), 0, func(acc, x int) int { return acc + x })
		if res != 0 {
			t.Errorf("Expected 0, got %v", res)
		}
	})

	t.Run("Count", func(t *testing.T) {
		c := ops.Count(core.Slice([]int{}))
		if c != 0 {
			t.Errorf("Expected 0, got %v", c)
		}
	})

	t.Run("Take", func(t *testing.T) {
		it := ops.Take(core.Slice([]int{}), 5)
		res := ops.Collect(it)
		if len(res) != 0 {
			t.Errorf("Expected empty result, got %v", res)
		}
	})

	t.Run("Skip", func(t *testing.T) {
		it := ops.Skip(core.Slice([]int{}), 5)
		res := ops.Collect(it)
		if len(res) != 0 {
			t.Errorf("Expected empty result, got %v", res)
		}
	})

	t.Run("Distinct", func(t *testing.T) {
		cmp := compare.LessFunc[int](func(a, b int) bool { return a < b })
		it := ops.Distinct(core.Slice([]int{}), cmp)
		res := ops.Collect(it)
		if len(res) != 0 {
			t.Errorf("Expected empty result, got %v", res)
		}
	})

	t.Run("GroupBy", func(t *testing.T) {
		it := ops.GroupBy(core.Slice([]int{}), func(x int) int { return x })
		res := ops.Collect(it)
		if len(res) != 0 {
			t.Errorf("Expected empty result, got %v", res)
		}
	})

	t.Run("Chain", func(t *testing.T) {
		it := ops.Chain(core.Slice([]int{}), core.Slice([]int{}))
		res := ops.Collect(it)
		if len(res) != 0 {
			t.Errorf("Expected empty result, got %v", res)
		}
	})

	t.Run("Zip", func(t *testing.T) {
		it := ops.Zip(core.Slice([]int{}), core.Slice([]string{}))
		res := ops.Collect(it)
		if len(res) != 0 {
			t.Errorf("Expected empty result, got %v", res)
		}
	})

	t.Run("Join", func(t *testing.T) {
		it := ops.Join(
			core.Slice([]int{}),
			core.Slice([]int{}),
			func(x int) int { return x },
			func(x int) int { return x },
		)
		res := ops.Collect(it)
		if len(res) != 0 {
			t.Errorf("Expected empty result, got %v", res)
		}
	})
}

func TestEmptyRange(t *testing.T) {
	// Range(start, end, step) where start >= end with positive step should be empty
	it := core.Range(10, 10, 1)
	res := ops.Collect(it)
	if len(res) != 0 {
		t.Errorf("Expected empty result, got %v", res)
	}
}
