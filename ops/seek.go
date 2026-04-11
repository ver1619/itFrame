package ops

import "github.com/ver1619/itFrame/core"

// SeekIterator skips elements until a predicate is satisfied, then iterates normally.
type SeekIterator[T any] struct {
	it     core.Iterator[T]
	pred   func(T) bool
	seeked bool
}

// Seek creates an iterator that skips elements until pred returns true.
// The first matching element is returned, and subsequent iteration continues normally.
func Seek[T any](
	it core.Iterator[T],
	pred func(T) bool,
) core.Iterator[T] {
	return &SeekIterator[T]{it: it, pred: pred}
}

// Next returns the next element after the seek point, or (zero, false) when exhausted.
func (s *SeekIterator[T]) Next() (T, bool) {
	if !s.seeked {
		for {
			val, ok := s.it.Next()
			if !ok {
				var zero T
				return zero, false
			}
			if s.pred(val) {
				s.seeked = true
				return val, true
			}
		}
	}

	return s.it.Next()
}
