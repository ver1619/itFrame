package ops

import "github.com/ver1619/itFrame/core"

// DropWhileIterator skips elements while the predicate returns true,
// then yields all remaining elements.
type DropWhileIterator[T any] struct {
	it      core.Iterator[T]
	pred    func(T) bool
	dropped bool
}

// DropWhile creates a DropWhileIterator that skips elements while pred is satisfied.
func DropWhile[T any](it core.Iterator[T], pred func(T) bool) core.Iterator[T] {
	return &DropWhileIterator[T]{it: it, pred: pred}
}

// Next returns the next element after the drop phase, or (zero, false) when exhausted.
func (d *DropWhileIterator[T]) Next() (T, bool) {
	if !d.dropped {
		for {
			val, ok := d.it.Next()
			if !ok {
				var zero T
				return zero, false
			}
			if !d.pred(val) {
				d.dropped = true
				return val, true
			}
		}
	}
	return d.it.Next()
}
