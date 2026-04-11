package ops_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestFlatMap_Basic(t *testing.T) {
	it := ops.FlatMap(
		core.Slice([]int{1, 2, 3}),
		func(v int) []int {
			return []int{v, v * 10}
		},
	)

	result := ops.Collect(it)
	expected := []int{1, 10, 2, 20, 3, 30}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestFlatMap_EmptyInner(t *testing.T) {
	it := ops.FlatMap(
		core.Slice([]int{1, 2, 3}),
		func(v int) []int {
			if v == 2 {
				return nil
			}
			return []int{v}
		},
	)

	result := ops.Collect(it)
	expected := []int{1, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
