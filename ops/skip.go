package ops

import "github.com/ver1619/itFrame/core"

// SkipIterator skips the first n elements, then yields the rest.
type SkipIterator[T any] struct {
	it      core.Iterator[T]
	n       int
	skipped bool
}

// Skip creates a SkipIterator that discards the first n elements.
func Skip[T any](it core.Iterator[T], n int) core.Iterator[T] {
	return &SkipIterator[T]{it: it, n: n}
}

// Next returns the next element after skipping, or (zero, false) when exhausted.
func (s *SkipIterator[T]) Next() (T, bool) {
	if !s.skipped {
		for i := 0; i < s.n; i++ {
			_, ok := s.it.Next()
			if !ok {
				var zero T
				return zero, false
			}
		}
		s.skipped = true
	}
	return s.it.Next()
}
