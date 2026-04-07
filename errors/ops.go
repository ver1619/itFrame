// Map (error-aware)
package errors

import "github.com/ver1619/itFrame/core"

type MapIterator[A, B any] struct {
	it core.Iterator[Result[A]]
	fn func(A) B
}

func Map[A, B any](
	it core.Iterator[Result[A]],
	fn func(A) B,
) core.Iterator[Result[B]] {
	return &MapIterator[A, B]{it: it, fn: fn}
}

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

// Filter (error-aware)

type FilterIterator[T any] struct {
	it   core.Iterator[Result[T]]
	pred func(T) bool
}

func Filter[T any](
	it core.Iterator[Result[T]],
	pred func(T) bool,
) core.Iterator[Result[T]] {
	return &FilterIterator[T]{it: it, pred: pred}
}

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

/*
- errors propagate automatically
- normal values are transformed
- operations do not suppress errors
*/
