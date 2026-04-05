package advanced_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/advanced"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestZip_Basic(t *testing.T) {
	it := advanced.Zip(
		core.Slice([]int{1, 2, 3}),
		core.Slice([]string{"a", "b", "c"}),
	)

	result := ops.Collect(it)

	expected := []advanced.Pair[int, string]{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestZip_Uneven(t *testing.T) {
	it := advanced.Zip(
		core.Slice([]int{1, 2, 3}),
		core.Slice([]string{"a", "b"}),
	)

	result := ops.Collect(it)

	expected := []advanced.Pair[int, string]{
		{1, "a"},
		{2, "b"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestZip_Empty(t *testing.T) {
	it := advanced.Zip(
		core.Slice([]int{}),
		core.Slice([]string{"a"}),
	)

	result := ops.Collect(it)

	if len(result) != 0 {
		t.Fatal("expected empty result")
	}
}
