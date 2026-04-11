package edge_tests

import (
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestNilHandling(t *testing.T) {
	t.Run("NilSlice", func(t *testing.T) {
		var s []int = nil
		it := core.Slice(s)
		_, ok := it.Next()
		if ok {
			t.Errorf("Expected nil slice iterator to be empty")
		}
	})

	t.Run("FlatMapReturningNil", func(t *testing.T) {
		it := ops.FlatMap(core.Slice([]int{1, 2, 3}), func(x int) []int {
			if x == 2 {
				return nil
			}
			return []int{x}
		})
		
		res := ops.Collect(it)
		if len(res) != 2 || res[0] != 1 || res[1] != 3 {
			t.Errorf("Expected [1 3], got %v", res)
		}
	})

	t.Run("MapReturningNilForInterface", func(t *testing.T) {
		// If T is an interface/pointer, Map returning nil should be valid
		it := ops.Map(core.Slice([]int{1}), func(x int) *int {
			return nil
		})
		v, ok := it.Next()
		if !ok || v != nil {
			t.Errorf("Expected ok=true and v=nil")
		}
	})
}
