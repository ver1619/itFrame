package ops_test

import (
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

type Result struct {
	Key int
	Sum int
}

func TestAggregate_Sum(t *testing.T) {
	grouped := ops.GroupBy(
		core.Slice([]int{1, 2, 3, 4}),
		func(x int) int { return x % 2 },
	)

	it := ops.Aggregate(
		grouped,
		func(g ops.Group[int, int]) Result {
			sum := 0
			for _, v := range g.Items {
				sum += v
			}
			return Result{g.Key, sum}
		},
	)

	var result []Result
	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		result = append(result, v)
	}

	if len(result) != 2 {
		t.Fatal("expected 2 groups")
	}
}
