package advanced_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/advanced"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func less(a, b int) bool { return a < b }

func TestMerge_Basic(t *testing.T) {
	it := advanced.Merge(
		core.Slice([]int{1, 3, 5}),
		core.Slice([]int{2, 4, 6}),
		less,
	)

	result := ops.Collect(it)

	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestMerge_DuplicatesStable(t *testing.T) {
	it := advanced.Merge(
		core.Slice([]int{1, 3, 5}),
		core.Slice([]int{1, 3, 4}),
		less,
	)

	result := ops.Collect(it)

	expected := []int{1, 1, 3, 3, 4, 5}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestMerge_Uneven(t *testing.T) {
	it := advanced.Merge(
		core.Slice([]int{1, 3}),
		core.Slice([]int{2, 4, 6}),
		less,
	)

	result := ops.Collect(it)

	expected := []int{1, 2, 3, 4, 6}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestMerge_OneEmpty(t *testing.T) {
	it := advanced.Merge(
		core.Slice([]int{}),
		core.Slice([]int{1, 2}),
		less,
	)

	result := ops.Collect(it)

	expected := []int{1, 2}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
