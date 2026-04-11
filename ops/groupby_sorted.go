package ops

import "github.com/ver1619/itFrame/core"

// GroupBySortedIterator yields groups lazily from a pre-sorted iterator.
// Unlike GroupBy, it does not buffer the entire input.
type GroupBySortedIterator[T any, K comparable] struct {
	it    core.Iterator[T]
	keyFn func(T) K

	buf     []T
	currKey K
	hasBuf  bool
	done    bool
}

// GroupBySorted creates a streaming GroupBy iterator for pre-sorted input.
// Elements must be sorted by the key function. Groups are emitted lazily
// as soon as the key changes, using O(group-size) memory instead of O(n).
func GroupBySorted[T any, K comparable](
	it core.Iterator[T],
	keyFn func(T) K,
) core.Iterator[Group[K, T]] {
	return &GroupBySortedIterator[T, K]{it: it, keyFn: keyFn}
}

// Next returns the next group, or (zero, false) when exhausted.
func (g *GroupBySortedIterator[T, K]) Next() (Group[K, T], bool) {
	if g.done {
		var zero Group[K, T]
		return zero, false
	}

	// seed with buffered element or fetch first
	if !g.hasBuf {
		v, ok := g.it.Next()
		if !ok {
			g.done = true
			var zero Group[K, T]
			return zero, false
		}
		g.currKey = g.keyFn(v)
		g.buf = []T{v}
		g.hasBuf = true
	}

	// collect all elements with the same key
	for {
		v, ok := g.it.Next()
		if !ok {
			g.done = true
			group := Group[K, T]{Key: g.currKey, Items: g.buf}
			g.buf = nil
			g.hasBuf = false
			return group, true
		}

		k := g.keyFn(v)
		if k != g.currKey {
			group := Group[K, T]{Key: g.currKey, Items: g.buf}
			g.currKey = k
			g.buf = []T{v}
			return group, true
		}

		g.buf = append(g.buf, v)
	}
}
