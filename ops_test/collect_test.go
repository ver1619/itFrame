package ops_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestCollect_Basic(t *testing.T) {
	it := core.Range(0, 5, 1)

	result := ops.Collect(it)

	expected := []int{0, 1, 2, 3, 4}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestCollect_Empty(t *testing.T) {
	it := core.Slice([]int{})

	result := ops.Collect(it)

	if len(result) != 0 {
		t.Fatal("expected empty slice")
	}
}
