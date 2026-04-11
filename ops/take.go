package ops

import "github.com/ver1619/itFrame/core"

// TakeIterator yields at most n elements from the underlying iterator.
type TakeIterator[T any] struct {
	it        core.Iterator[T]
	remaining int
}

// Take creates a TakeIterator that yields at most n elements.
func Take[T any](it core.Iterator[T], n int) core.Iterator[T] {
	return &TakeIterator[T]{it: it, remaining: n}
}

// Next returns the next element if the limit has not been reached, or (zero, false) otherwise.
func (t *TakeIterator[T]) Next() (T, bool) {
	if t.remaining <= 0 {
		var zero T
		return zero, false
	}
	t.remaining--
	return t.it.Next()
}
