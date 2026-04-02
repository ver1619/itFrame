package ops

import "github.com/ver1619/itFrame/core"

type FilterIterator[T any] struct {
	it   core.Iterator[T]
	pred func(T) bool
}

func Filter[T any](it core.Iterator[T], pred func(T) bool) *FilterIterator[T] {
	return &FilterIterator[T]{
		it:   it,
		pred: pred,
	}
}

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

/*FilterIterator returns only values that satisfy a condition.
Filter(it, pred) creates a new iterator using a predicate function.

Each call to Next():
- pulls values from underlying iterator
- skips values that don’t match pred
- returns first matching value
*/
