package ops

import "github.com/ver1619/itFrame/core"

// FilterIterator yields only elements that satisfy a predicate.
type FilterIterator[T any] struct {
	it   core.Iterator[T]
	pred func(T) bool
}

// Filter creates a FilterIterator that keeps elements matching pred.
func Filter[T any](it core.Iterator[T], pred func(T) bool) core.Iterator[T] {
	return &FilterIterator[T]{
		it:   it,
		pred: pred,
	}
}

// Next returns the next element satisfying the predicate, or (zero, false) when exhausted.
func (f *FilterIterator[T]) Next() (T, bool) {
	for {
		val, ok := f.it.Next()
		if !ok {
			var zero T
			return zero, false
		}
		if f.pred(val) {
			return val, true
		}
	}
}
