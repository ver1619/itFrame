package ops_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/compare"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestDistinct_Basic(t *testing.T) {
	it := ops.Distinct(
		core.Slice([]int{1, 1, 2, 3, 3, 3, 4}),
		compare.LessFunc[int](func(a, b int) bool { return a < b }),
	)

	result := ops.Collect(it)
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestDistinct_AllSame(t *testing.T) {
	it := ops.Distinct(
		core.Slice([]int{5, 5, 5}),
		compare.LessFunc[int](func(a, b int) bool { return a < b }),
	)

	result := ops.Collect(it)
	expected := []int{5}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestDistinct_NoDuplicates(t *testing.T) {
	it := ops.Distinct(
		core.Slice([]int{1, 2, 3}),
		compare.LessFunc[int](func(a, b int) bool { return a < b }),
	)

	result := ops.Collect(it)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestDistinct_Empty(t *testing.T) {
	it := ops.Distinct(
		core.Slice([]int{}),
		compare.LessFunc[int](func(a, b int) bool { return a < b }),
	)

	result := ops.Collect(it)
	if len(result) != 0 {
		t.Fatalf("expected empty, got %v", result)
	}
}
