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
		func(x int) core.Iterator[int] {
			return core.Slice([]int{x, x * x})
		},
	)

	var result []int
	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		result = append(result, v)
	}

	expected := []int{1, 1, 2, 4, 3, 9}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestFlatMap_EmptyInner(t *testing.T) {
	it := ops.FlatMap(
		core.Slice([]int{1, 2}),
		func(x int) core.Iterator[int] {
			if x == 1 {
				return core.Slice([]int{})
			}
			return core.Slice([]int{2})
		},
	)

	var result []int
	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		result = append(result, v)
	}

	expected := []int{2}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
