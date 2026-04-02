package ops_test

import (
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestComposition_MapThenFilter(t *testing.T) {
	it := ops.Filter(
		ops.Map(
			core.Slice([]int{1, 2, 3, 4}),
			func(x int) int { return x * x },
		),
		func(x int) bool { return x%2 == 0 },
	)

	expected := []int{4, 16}

	for i, v := range expected {
		val, ok := it.Next()
		if !ok || val != v {
			t.Fatalf("expected %d at index %d", v, i)
		}
	}
}

func TestComposition_FilterThenMap(t *testing.T) {
	it := ops.Map(
		ops.Filter(
			core.Slice([]int{1, 2, 3, 4}),
			func(x int) bool { return x%2 == 0 },
		),
		func(x int) int { return x * 10 },
	)

	expected := []int{20, 40}

	for i, v := range expected {
		val, ok := it.Next()
		if !ok || val != v {
			t.Fatalf("expected %d at index %d", v, i)
		}
	}
}
