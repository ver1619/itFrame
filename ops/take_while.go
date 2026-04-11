package ops

import "github.com/ver1619/itFrame/core"

// TakeWhileIterator yields elements while the predicate returns true.
// Stops permanently on the first element that fails the predicate.
type TakeWhileIterator[T any] struct {
	it   core.Iterator[T]
	pred func(T) bool
	done bool
}

// TakeWhile creates a TakeWhileIterator that yields elements while pred is satisfied.
func TakeWhile[T any](it core.Iterator[T], pred func(T) bool) core.Iterator[T] {
	return &TakeWhileIterator[T]{it: it, pred: pred}
}

// Next returns the next element if the predicate holds, or (zero, false) once it fails.
func (t *TakeWhileIterator[T]) Next() (T, bool) {
	if t.done {
		var zero T
		return zero, false
	}

	val, ok := t.it.Next()
	if !ok {
		t.done = true
		var zero T
		return zero, false
	}

	if !t.pred(val) {
		t.done = true
		var zero T
		return zero, false
	}

	return val, true
}
