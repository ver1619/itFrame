package ops

import "github.com/ver1619/itFrame/core"

type AggregateIterator[K comparable, V any, R any] struct {
	it core.Iterator[Group[K, V]]
	fn func(Group[K, V]) R
}

func Aggregate[K comparable, V any, R any](
	it core.Iterator[Group[K, V]],
	fn func(Group[K, V]) R,
) core.Iterator[R] {
	return &AggregateIterator[K, V, R]{
		it: it,
		fn: fn,
	}
}

func (a *AggregateIterator[K, V, R]) Next() (R, bool) {
	g, ok := a.it.Next()
	if !ok {
		var zero R
		return zero, false
	}
	return a.fn(g), true
}

/*
- **Aggregate** applies a function to each group.
- Produces summarized results.
- Fully lazy over groups
*/
