package ops

import "github.com/ver1619/itFrame/core"

// ChainIterator concatenates two iterators sequentially.
// All elements from the first iterator are yielded before the second begins.
type ChainIterator[T any] struct {
	it1   core.Iterator[T]
	it2   core.Iterator[T]
	first bool
}

// Chain creates a ChainIterator that yields all elements from it1, then all from it2.
func Chain[T any](it1, it2 core.Iterator[T]) core.Iterator[T] {
	return &ChainIterator[T]{it1: it1, it2: it2, first: true}
}

// Next returns the next element from the chained sequence, or (zero, false) when both are exhausted.
func (c *ChainIterator[T]) Next() (T, bool) {
	if c.first {
		val, ok := c.it1.Next()
		if ok {
			return val, true
		}
		c.first = false
	}
	return c.it2.Next()
}
