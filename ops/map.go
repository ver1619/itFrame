// Package ops provides composable iterator operations including transformations,
// filters, terminal operations, and relational joins.
package ops

import "github.com/ver1619/itFrame/core"

// MapIterator applies a transformation function to each element.
type MapIterator[A, B any] struct {
	it core.Iterator[A]
	fn func(A) B
}

// Map creates a MapIterator that transforms each element using fn.
func Map[A, B any](it core.Iterator[A], fn func(A) B) core.Iterator[B] {
	return &MapIterator[A, B]{
		it: it,
		fn: fn,
	}
}

// Next returns the next transformed element, or (zero, false) when exhausted.
func (m *MapIterator[A, B]) Next() (B, bool) {
	val, ok := m.it.Next()
	if !ok {
		var zero B
		return zero, false
	}
	return m.fn(val), true
}
