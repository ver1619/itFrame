package ops_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestJoin_Basic(t *testing.T) {
	type A struct{ ID int }
	type B struct{ ID int }

	it := ops.Join(
		core.Slice([]A{{1}, {2}}),
		core.Slice([]B{{1}, {1}, {3}}),
		func(a A) int { return a.ID },
		func(b B) int { return b.ID },
	)

	result := ops.Collect(it)

	expected := []ops.Pair[A, B]{
		{First: A{1}, Second: B{1}},
		{First: A{1}, Second: B{1}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
