package ops

import "github.com/ver1619/itFrame/core"

// FlatMapIterIterator maps each element to an iterator and flattens the results.
type FlatMapIterIterator[A, B any] struct {
	outer core.Iterator[A]
	fn    func(A) core.Iterator[B]

	inner core.Iterator[B]
}

// FlatMapIter creates a FlatMapIterIterator that expands each element into an iterator and flattens.
// Nil inner iterators are skipped safely.
func FlatMapIter[A, B any](
	it core.Iterator[A],
	fn func(A) core.Iterator[B],
) core.Iterator[B] {
	return &FlatMapIterIterator[A, B]{
		outer: it,
		fn:    fn,
	}
}

// Next returns the next flattened element, or (zero, false) when exhausted.
func (f *FlatMapIterIterator[A, B]) Next() (B, bool) {
	for {
		if f.inner != nil {
			val, ok := f.inner.Next()
			if ok {
				return val, true
			}
			f.inner = nil
		}

		outerVal, ok := f.outer.Next()
		if !ok {
			var zero B
			return zero, false
		}

		f.inner = f.fn(outerVal)

		if f.inner == nil {
			continue
		}
	}
}
