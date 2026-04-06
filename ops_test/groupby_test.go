package ops_test

import (
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestGroupBy_Count(t *testing.T) {
	it := ops.GroupBy(
		core.Slice([]int{1, 2, 3, 4}),
		func(x int) int { return x % 2 },
	)

	counts := make(map[int]int)

	for {
		g, ok := it.Next()
		if !ok {
			break
		}
		counts[g.Key] = len(g.Items)
	}

	if counts[0] != 2 || counts[1] != 2 {
		t.Fatal("incorrect grouping")
	}
}
