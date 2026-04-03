package ops_test

import (
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestCount_Basic(t *testing.T) {
	it := core.Range(0, 5, 1)

	c := ops.Count(it)

	if c != 5 {
		t.Fatalf("expected 5, got %d", c)
	}
}

func TestCount_WithFilter(t *testing.T) {
	it := ops.Filter(
		core.Range(0, 10, 1),
		func(x int) bool { return x%2 == 0 },
	)

	c := ops.Count(it)

	if c != 5 {
		t.Fatalf("expected 5, got %d", c)
	}
}
