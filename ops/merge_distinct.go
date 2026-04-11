package ops

import (
	"github.com/ver1619/itFrame/compare"
	"github.com/ver1619/itFrame/core"
)

// MergeDistinctIterator merges two sorted iterators while removing duplicates.
type MergeDistinctIterator[T any] struct {
	it1   core.Iterator[T]
	it2   core.Iterator[T]
	cmp   compare.Comparator[T]
	v1    T
	v2    T
	ok1   bool
	ok2   bool
	initd bool

	last    T
	hasLast bool
}

// MergeDistinct creates an iterator that merges two sorted iterators,
// emitting each distinct value only once.
func MergeDistinct[T any](
	it1, it2 core.Iterator[T],
	cmp compare.Comparator[T],
) core.Iterator[T] {
	return &MergeDistinctIterator[T]{
		it1: it1,
		it2: it2,
		cmp: cmp,
	}
}

func (m *MergeDistinctIterator[T]) init() {
	if m.initd {
		return
	}
	m.v1, m.ok1 = m.it1.Next()
	m.v2, m.ok2 = m.it2.Next()
	m.initd = true
}

// Next returns the next distinct element in sorted order, or (zero, false) when exhausted.
func (m *MergeDistinctIterator[T]) Next() (T, bool) {
	m.init()

	for {
		if !m.ok1 && !m.ok2 {
			var zero T
			return zero, false
		}

		var val T

		if !m.ok1 {
			val = m.v2
			m.v2, m.ok2 = m.it2.Next()
		} else if !m.ok2 {
			val = m.v1
			m.v1, m.ok1 = m.it1.Next()
		} else if m.cmp.Less(m.v1, m.v2) {
			val = m.v1
			m.v1, m.ok1 = m.it1.Next()
		} else if m.cmp.Less(m.v2, m.v1) {
			val = m.v2
			m.v2, m.ok2 = m.it2.Next()
		} else {
			val = m.v1
			m.v1, m.ok1 = m.it1.Next()
			m.v2, m.ok2 = m.it2.Next()
		}

		if m.hasLast && compare.Equal(m.cmp, val, m.last) {
			continue
		}

		m.last = val
		m.hasLast = true
		return val, true
	}
}
