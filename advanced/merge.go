package advanced

import (
	"github.com/ver1619/itFrame/compare"
	"github.com/ver1619/itFrame/core"
)

type MergeIterator[T any] struct {
	it1   core.Iterator[T]
	it2   core.Iterator[T]
	cmp   compare.Comparator[T]
	v1    T
	v2    T
	ok1   bool
	ok2   bool
	initd bool
}

func Merge[T any](
	it1, it2 core.Iterator[T],
	cmp compare.Comparator[T],
) *MergeIterator[T] {
	return &MergeIterator[T]{
		it1: it1,
		it2: it2,
		cmp: cmp,
	}
}

func (m *MergeIterator[T]) init() {
	if m.initd {
		return
	}
	m.v1, m.ok1 = m.it1.Next()
	m.v2, m.ok2 = m.it2.Next()
	m.initd = true
}

func (m *MergeIterator[T]) Next() (T, bool) {
	m.init()

	if !m.ok1 && !m.ok2 {
		var zero T
		return zero, false
	}

	if !m.ok1 {
		val := m.v2
		m.v2, m.ok2 = m.it2.Next()
		return val, true
	}

	if !m.ok2 {
		val := m.v1
		m.v1, m.ok1 = m.it1.Next()
		return val, true
	}

	if m.cmp.Less(m.v1, m.v2) {
		val := m.v1
		m.v1, m.ok1 = m.it1.Next()
		return val, true
	}

	if m.cmp.Less(m.v2, m.v1) {
		val := m.v2
		m.v2, m.ok2 = m.it2.Next()
		return val, true
	}

	// equal → stable: prefer it1
	val := m.v1
	m.v1, m.ok1 = m.it1.Next()
	return val, true
}

/*

- **MergeIterator** combines two sorted iterators into one sorted sequence.
- Merge(it1, it2, less) takes a comparator function to define ordering.
- Each call to Next() returns the smallest next element from both iterators.
- Works lazily — does not load all data into memory.
- stable merge (prefer it1 on equality). When values are equal, MergeIterator prefers elements from the first iterator.

*/

/*
v0.6.0
- **Merge** now accepts a Comparator instead of a raw function.
- Ordering logic is centralized and reusable across features.
- Behavior remains the same: stable, lazy merge of two sorted iterators.
- Equality is derived via comparator → consistent handling of duplicates.
*/
