package ops_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestZipLongest_Uneven(t *testing.T) {
	it := ops.ZipLongest(
		core.Slice([]int{1, 2, 3}),
		core.Slice([]string{"a", "b"}),
	)

	result := ops.Collect(it)

	expected := []ops.Pair[int, string]{
		{First: 1, Second: "a"},
		{First: 2, Second: "b"},
		{First: 3, Second: ""},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
