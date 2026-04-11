package ops

import "github.com/ver1619/itFrame/core"

// Group holds a key and its associated items produced by GroupBy.
type Group[K comparable, V any] struct {
	Key   K
	Items []V
}

// GroupByIterator yields groups of elements sharing the same key.
type GroupByIterator[T any, K comparable] struct {
	groups []Group[K, T]
	idx    int
}

// GroupBy consumes the iterator and groups elements by the key returned from keyFn.
// The entire input is buffered internally. Group order is not guaranteed.
func GroupBy[T any, K comparable](
	it core.Iterator[T],
	keyFn func(T) K,
) core.Iterator[Group[K, T]] {

	groupMap := make(map[K][]T)

	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		k := keyFn(v)
		groupMap[k] = append(groupMap[k], v)
	}

	groups := make([]Group[K, T], 0, len(groupMap))
	for k, items := range groupMap {
		groups = append(groups, Group[K, T]{
			Key:   k,
			Items: items,
		})
	}

	return &GroupByIterator[T, K]{groups: groups}
}

// Next returns the next group, or (zero, false) when exhausted.
func (g *GroupByIterator[T, K]) Next() (Group[K, T], bool) {
	if g.idx >= len(g.groups) {
		var zero Group[K, T]
		return zero, false
	}

	val := g.groups[g.idx]
	g.idx++
	return val, true
}
