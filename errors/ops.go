package errors

import "github.com/ver1619/itFrame/core"

// MapIterator applies a transformation to valid values, propagating errors unchanged.
type MapIterator[A, B any] struct {
	it core.Iterator[Result[A]]
	fn func(A) B
}

// Map creates an error-aware MapIterator. Errors propagate automatically.
func Map[A, B any](
	it core.Iterator[Result[A]],
	fn func(A) B,
) core.Iterator[Result[B]] {
	return &MapIterator[A, B]{it: it, fn: fn}
}

// Next returns the next transformed result, or (zero, false) when exhausted.
func (m *MapIterator[A, B]) Next() (Result[B], bool) {
	r, ok := m.it.Next()
	if !ok {
		var zero Result[B]
		return zero, false
	}

	if r.Err != nil {
		return ErrResult[B](r.Err), true
	}

	return Ok(m.fn(r.Value)), true
}

// FilterIterator yields valid values that satisfy a predicate, propagating errors unchanged.
type FilterIterator[T any] struct {
	it   core.Iterator[Result[T]]
	pred func(T) bool
}

// Filter creates an error-aware FilterIterator. Errors propagate automatically.
func Filter[T any](
	it core.Iterator[Result[T]],
	pred func(T) bool,
) core.Iterator[Result[T]] {
	return &FilterIterator[T]{it: it, pred: pred}
}

// Next returns the next matching result, or (zero, false) when exhausted.
func (f *FilterIterator[T]) Next() (Result[T], bool) {
	for {
		r, ok := f.it.Next()
		if !ok {
			var zero Result[T]
			return zero, false
		}

		if r.Err != nil {
			return r, true
		}

		if f.pred(r.Value) {
			return r, true
		}
	}
}
