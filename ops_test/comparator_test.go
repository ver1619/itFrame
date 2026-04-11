package ops_test

import (
	"testing"

	"github.com/ver1619/itFrame/compare"
)

func TestComparator_Equal(t *testing.T) {
	cmp := compare.LessFunc[int](func(a, b int) bool { return a < b })

	if !compare.Equal(cmp, 5, 5) {
		t.Fatal("expected equal")
	}
}

func TestComparator_Reverse(t *testing.T) {
	base := compare.LessFunc[int](func(a, b int) bool { return a < b })
	rev := compare.Reverse[int]{Cmp: base}

	if !rev.Less(5, 3) {
		t.Fatal("expected reverse ordering")
	}
}
