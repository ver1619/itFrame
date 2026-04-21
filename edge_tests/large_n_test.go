package edge_tests

import (
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestLargeN(t *testing.T) {
	// 10 million elements
	const n = 10_000_000

	t.Run("RangeLaziness", func(t *testing.T) {
		// Range should be instantaneous even for 10M elements
		it := core.Range(0, n, 1)

		// Take only first 3
		res := ops.Take(it, 3)
		slice := ops.Collect(res)

		if len(slice) != 3 {
			t.Errorf("expected 3 elements, got %d", len(slice))
		}
	})

	t.Run("CountLarge", func(t *testing.T) {
		it := core.Range(0, n, 1)
		c := ops.Count(it)
		if c != n {
			t.Errorf("expected %d, got %d", n, c)
		}
	})

	t.Run("DeeplyNestedMap", func(t *testing.T) {
		// Verify that a deep chain doesn't hit stack limits (shouldn't, it's pull-based)
		var it core.Iterator[int] = core.Range(0, 1000, 1)
		for i := 0; i < 100; i++ {
			it = ops.Map(it, func(x int) int { return x + 1 })
		}

		v, ok := it.Next()
		if !ok || v != 100 {
			t.Errorf("expected 100, got %v", v)
		}
	})
}
