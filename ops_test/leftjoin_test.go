package ops_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/advanced"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestLeftJoin_Basic(t *testing.T) {
	type A struct{ ID int }
	type B struct{ ID int }

	it := ops.LeftJoin(
		core.Slice([]A{{1}, {2}}),
		core.Slice([]B{{1}}),
		func(a A) int { return a.ID },
		func(b B) int { return b.ID },
	)

	result := ops.Collect(it)

	expected := []advanced.Pair[A, B]{
		{First: A{1}, Second: B{1}},
		{First: A{2}, Second: B{}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
