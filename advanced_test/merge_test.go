package advanced_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/advanced"
	"github.com/ver1619/itFrame/compare"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func less(a, b int) bool { return a < b }

func TestMerge_Basic(t *testing.T) {
	cmp := compare.LessFunc[int](less)

	it := advanced.Merge(
		core.Slice([]int{1, 3, 5}),
		core.Slice([]int{2, 4, 6}),
		cmp,
	)

	result := ops.Collect(it)

	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestMerge_DuplicatesStable(t *testing.T) {
	cmp := compare.LessFunc[int](less)

	it := advanced.Merge(
		core.Slice([]int{1, 3, 5}),
		core.Slice([]int{1, 3, 4}),
		cmp,
	)

	result := ops.Collect(it)

	// stable: first iterator preferred on equality
	expected := []int{1, 1, 3, 3, 4, 5}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestMerge_Uneven(t *testing.T) {
	cmp := compare.LessFunc[int](less)

	it := advanced.Merge(
		core.Slice([]int{1, 3}),
		core.Slice([]int{2, 4, 6}),
		cmp,
	)

	result := ops.Collect(it)

	expected := []int{1, 2, 3, 4, 6}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestMerge_OneEmpty(t *testing.T) {
	cmp := compare.LessFunc[int](less)

	it := advanced.Merge(
		core.Slice([]int{}),
		core.Slice([]int{1, 2}),
		cmp,
	)

	result := ops.Collect(it)

	expected := []int{1, 2}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestMerge_BothEmpty(t *testing.T) {
	cmp := compare.LessFunc[int](less)

	it := advanced.Merge(
		core.Slice([]int{}),
		core.Slice([]int{}),
		cmp,
	)

	result := ops.Collect(it)

	if len(result) != 0 {
		t.Fatal("expected empty result")
	}
}
