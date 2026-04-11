package ops_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestZip_Basic(t *testing.T) {
	it := ops.Zip(
		core.Slice([]int{1, 2, 3}),
		core.Slice([]string{"a", "b", "c"}),
	)

	result := ops.Collect(it)

	expected := []ops.Pair[int, string]{
		{First: 1, Second: "a"},
		{First: 2, Second: "b"},
		{First: 3, Second: "c"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestZip_Uneven(t *testing.T) {
	it := ops.Zip(
		core.Slice([]int{1, 2, 3}),
		core.Slice([]string{"a", "b"}),
	)

	result := ops.Collect(it)

	expected := []ops.Pair[int, string]{
		{First: 1, Second: "a"},
		{First: 2, Second: "b"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestZip_Empty(t *testing.T) {
	it := ops.Zip(
		core.Slice([]int{}),
		core.Slice([]string{"a"}),
	)

	result := ops.Collect(it)

	if len(result) != 0 {
		t.Fatal("expected empty result")
	}
}
