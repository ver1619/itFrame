package edge_tests

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/compare"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestLengthImbalance(t *testing.T) {
	t.Run("ZipShortLeft", func(t *testing.T) {
		left := core.Slice([]int{1, 2})
		right := core.Slice([]string{"a", "b", "c", "d"})
		
		it := ops.Zip(left, right)
		res := ops.Collect(it)
		
		if len(res) != 2 {
			t.Errorf("Expected 2 pairs, got %d", len(res))
		}
	})

	t.Run("ZipShortRight", func(t *testing.T) {
		left := core.Slice([]int{1, 2, 3, 4})
		right := core.Slice([]string{"a", "b"})
		
		it := ops.Zip(left, right)
		res := ops.Collect(it)
		
		if len(res) != 2 {
			t.Errorf("Expected 2 pairs, got %d", len(res))
		}
	})

	t.Run("ZipLongestShortLeft", func(t *testing.T) {
		left := core.Slice([]int{1, 2})
		right := core.Slice([]string{"a", "b", "c"})
		
		it := ops.ZipLongest(left, right)
		res := ops.Collect(it)
		
		if len(res) != 3 {
			t.Errorf("Expected 3 pairs, got %d", len(res))
		}
		if res[2].First != 0 || res[2].Second != "c" {
			t.Errorf("Expected {0, c}, got %v", res[2])
		}
	})

	t.Run("MergeShortLeft", func(t *testing.T) {
		cmp := compare.LessFunc[int](func(a, b int) bool { return a < b })
		left := core.Slice([]int{1, 10})
		right := core.Slice([]int{2, 3, 4})
		
		it := ops.Merge(left, right, cmp)
		res := ops.Collect(it)
		
		expected := []int{1, 2, 3, 4, 10}
		if !reflect.DeepEqual(res, expected) {
			t.Errorf("Expected %v, got %v", expected, res)
		}
	})
	
	t.Run("MergeShortRight", func(t *testing.T) {
		cmp := compare.LessFunc[int](func(a, b int) bool { return a < b })
		left := core.Slice([]int{2, 3, 4})
		right := core.Slice([]int{1, 10})
		
		it := ops.Merge(left, right, cmp)
		res := ops.Collect(it)
		
		expected := []int{1, 2, 3, 4, 10}
		if !reflect.DeepEqual(res, expected) {
			t.Errorf("Expected %v, got %v", expected, res)
		}
	})
}
