package errors

import "github.com/ver1619/itFrame/core"

type flatMapIterator[A, B any] struct {
	outer core.Iterator[Result[A]]
	fn    func(A) core.Iterator[Result[B]]

	inner core.Iterator[Result[B]]
}

func (s Stream[A]) FlatMap(fn func(A) core.Iterator[Result[A]]) Stream[A] {
	return Stream[A]{
		it: &flatMapIterator[A, A]{
			outer: s.it,
			fn:    fn,
		},
	}
}

func (f *flatMapIterator[A, B]) Next() (Result[B], bool) {
	for {
		if f.inner != nil {
			v, ok := f.inner.Next()
			if ok {
				return v, true
			}
			f.inner = nil
		}

		r, ok := f.outer.Next()
		if !ok {
			var zero Result[B]
			return zero, false
		}

		if r.Err != nil {
			return ErrResult[B](r.Err), true
		}

		f.inner = f.fn(r.Value)
		if f.inner == nil {
			continue
		}
	}
}

/*
- propagates errors immediately
- expand valid values
*/
