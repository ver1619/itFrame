package advanced_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/advanced"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestZipLongest_Uneven(t *testing.T) {
	it := advanced.ZipLongest(
		core.Slice([]int{1, 2, 3}),
		core.Slice([]string{"a", "b"}),
	)

	result := ops.Collect(it)

	expected := []advanced.Pair[int, string]{
		{1, "a"},
		{2, "b"},
		{3, ""},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
