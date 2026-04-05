package advanced_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/advanced"
	"github.com/ver1619/itFrame/compare"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestMergeDistinct_Basic(t *testing.T) {
	cmp := compare.LessFunc[int](less)

	it := advanced.MergeDistinct(
		core.Slice([]int{1, 2, 3}),
		core.Slice([]int{2, 3, 4}),
		cmp,
	)

	result := ops.Collect(it)

	expected := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestMergeDistinct_DuplicatesInside(t *testing.T) {
	cmp := compare.LessFunc[int](less)

	it := advanced.MergeDistinct(
		core.Slice([]int{1, 1, 2}),
		core.Slice([]int{2, 2, 3}),
		cmp,
	)

	result := ops.Collect(it)

	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
