package ops_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestZipWith_Basic(t *testing.T) {
	it := ops.ZipWith(
		core.Slice([]int{1, 2, 3}),
		core.Slice([]int{4, 5, 6}),
		func(a, b int) int { return a + b },
	)

	result := ops.Collect(it)

	expected := []int{5, 7, 9}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
