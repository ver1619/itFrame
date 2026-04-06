package ops

import "github.com/ver1619/itFrame/core"

type Group[K comparable, V any] struct {
	Key   K
	Items []V
}

type GroupByIterator[T any, K comparable] struct {
	groups []Group[K, T]
	idx    int
}

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

	// convert map → slice
	groups := make([]Group[K, T], 0, len(groupMap))
	for k, items := range groupMap {
		groups = append(groups, Group[K, T]{
			Key:   k,
			Items: items,
		})
	}

	return &GroupByIterator[T, K]{groups: groups}
}

func (g *GroupByIterator[T, K]) Next() (Group[K, T], bool) {
	if g.idx >= len(g.groups) {
		var zero Group[K, T]
		return zero, false
	}

	val := g.groups[g.idx]
	g.idx++
	return val, true
}

/*
- **GroupBy** groups elements by a key.
- Entire input is buffered internally.
- Each group contains a key and its items.
- Order of groups is not guaranteed.
*/
