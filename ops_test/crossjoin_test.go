package ops_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/advanced"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestCrossJoin_Basic(t *testing.T) {
	it := ops.CrossJoin(
		core.Slice([]int{1, 2}),
		core.Slice([]string{"a", "b"}),
	)

	result := ops.Collect(it)

	expected := []advanced.Pair[int, string]{
		{First: 1, Second: "a"}, {First: 1, Second: "b"},
		{First: 2, Second: "a"}, {First: 2, Second: "b"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
