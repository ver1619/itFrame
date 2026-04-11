package ops

import "github.com/ver1619/itFrame/core"

// AggregateIterator applies a summarization function to each group.
type AggregateIterator[K comparable, V any, R any] struct {
	it core.Iterator[Group[K, V]]
	fn func(Group[K, V]) R
}

// Aggregate creates an iterator that applies fn to each group, producing summarized results.
func Aggregate[K comparable, V any, R any](
	it core.Iterator[Group[K, V]],
	fn func(Group[K, V]) R,
) core.Iterator[R] {
	return &AggregateIterator[K, V, R]{
		it: it,
		fn: fn,
	}
}

// Next returns the next aggregated result, or (zero, false) when exhausted.
func (a *AggregateIterator[K, V, R]) Next() (R, bool) {
	g, ok := a.it.Next()
	if !ok {
		var zero R
		return zero, false
	}
	return a.fn(g), true
}
