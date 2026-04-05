package advanced_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/advanced"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestSeek_Basic(t *testing.T) {
	it := advanced.Seek(
		core.Range(0, 10, 1),
		func(x int) bool { return x >= 5 },
	)

	result := ops.Collect(it)

	expected := []int{5, 6, 7, 8, 9}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestSeek_NoMatch(t *testing.T) {
	it := advanced.Seek(
		core.Range(0, 5, 1),
		func(x int) bool { return x > 10 },
	)

	result := ops.Collect(it)

	if len(result) != 0 {
		t.Fatal("expected empty result")
	}
}
