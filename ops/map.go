package ops

import "github.com/ver1619/itFrame/core"

type MapIterator[A, B any] struct {
	it core.Iterator[A]
	fn func(A) B
}

func Map[A, B any](it core.Iterator[A], fn func(A) B) *MapIterator[A, B] {
	return &MapIterator[A, B]{
		it: it,
		fn: fn,
	}
}

func (m *MapIterator[A, B]) Next() (B, bool) {
	val, ok := m.it.Next()
	if !ok {
		var zero B
		return zero, false
	}
	return m.fn(val), true
}
