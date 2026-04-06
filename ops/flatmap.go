package ops

import "github.com/ver1619/itFrame/core"

type FlatMapIterator[A, B any] struct {
	outer core.Iterator[A]
	fn    func(A) core.Iterator[B]

	inner core.Iterator[B]
}

func FlatMap[A, B any](
	it core.Iterator[A],
	fn func(A) core.Iterator[B],
) core.Iterator[B] {
	return &FlatMapIterator[A, B]{
		outer: it,
		fn:    fn,
	}
}

func (f *FlatMapIterator[A, B]) Next() (B, bool) {
	for {
		// if we have an active inner iterator
		if f.inner != nil {
			val, ok := f.inner.Next()
			if ok {
				return val, true
			}
			// inner exhausted → reset
			f.inner = nil
		}

		// get next outer element
		outerVal, ok := f.outer.Next()
		if !ok {
			var zero B
			return zero, false
		}

		// create new inner iterator
		f.inner = f.fn(outerVal)

		// if fn returns nil → skip safely
		if f.inner == nil {
			continue
		}
	}
}

/*
- **FlatMap** converts each element into a stream and flattens the result
- Maps each element to an iterator.
- One input can produce multiple outputs.
- Skips empty or nil inner iterators.
- lazy — values are generated on demand.
*/
