package advanced

import "github.com/ver1619/itFrame/core"

type SeekIterator[T any] struct {
	it     core.Iterator[T]
	pred   func(T) bool
	seeked bool
}

func Seek[T any](
	it core.Iterator[T],
	pred func(T) bool,
) *SeekIterator[T] {
	return &SeekIterator[T]{
		it:   it,
		pred: pred,
	}
}

func (s *SeekIterator[T]) Next() (T, bool) {
	// perform seek only once
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

/*
- **Seek** skips elements until predicate is true.
- First matching element is returned.
- After that, iteration continues normally.
- Works lazily (no preloading).
- Seek happens only once.
*/
