package ops

import (
	"github.com/ver1619/itFrame/compare"
	"github.com/ver1619/itFrame/core"
)

// MergeIterator combines two sorted iterators into one sorted sequence.
// When values are equal, elements from the first iterator are preferred (stable merge).
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

// Merge creates an iterator that merges two sorted iterators using the given comparator.
func Merge[T any](
	it1, it2 core.Iterator[T],
	cmp compare.Comparator[T],
) core.Iterator[T] {
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

// Next returns the next element in sorted order, or (zero, false) when both iterators are exhausted.
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

	val := m.v1
	m.v1, m.ok1 = m.it1.Next()
	return val, true
}
