package edge_tests

import (
	"testing"

	"github.com/ver1619/itFrame/core"
)

func TestIteratorDoubleConsumption(t *testing.T) {
	t.Run("SliceIterator", func(t *testing.T) {
		it := core.Slice([]int{1, 2})
		
		// First pass
		v1, ok1 := it.Next()
		if !ok1 || v1 != 1 {
			t.Errorf("Expected 1, got %v", v1)
		}
		v2, ok2 := it.Next()
		if !ok2 || v2 != 2 {
			t.Errorf("Expected 2, got %v", v2)
		}
		_, ok3 := it.Next()
		if ok3 {
			t.Errorf("Expected false after exhaustion")
		}

		// Double consumption - subsequent calls to Next should still return false
		_, ok4 := it.Next()
		if ok4 {
			t.Errorf("Expected false on second call after exhaustion")
		}
	})

	t.Run("RangeIterator", func(t *testing.T) {
		it := core.Range(0, 1, 1)
		v, ok := it.Next()
		if !ok || v != 0 {
			t.Errorf("Expected 0, got %v", v)
		}
		_, ok = it.Next()
		if ok {
			t.Errorf("Expected end")
		}
		_, ok = it.Next()
		if ok {
			t.Errorf("Expected end on repeat")
		}
	})
}

func TestPartialConsumption(t *testing.T) {
	// Verify that we can stop mid-iterator and it doesn't cause issues for the iterator object itself
	// (though usually you don't use it again, we should check it doesn't panic)
	it := core.Slice([]int{1, 2, 3, 4, 5})
	v, ok := it.Next()
	if !ok || v != 1 {
		t.Fatal("expected 1")
	}
	// Stop here. No cleanup needed for memory-based iterators, but good to have in suite.
}
