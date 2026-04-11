package ops

import "github.com/ver1619/itFrame/core"

// FlatMapIterator maps each element to a slice and flattens the results.
type FlatMapIterator[A, B any] struct {
	outer core.Iterator[A]
	fn    func(A) []B

	inner []B
	idx   int
}

// FlatMap creates a FlatMapIterator that expands each element into a slice and flattens.
// Use this instead of FlatMap when returning slices to avoid interface boxing allocations.
func FlatMap[A, B any](
	it core.Iterator[A],
	fn func(A) []B,
) core.Iterator[B] {
	return &FlatMapIterator[A, B]{
		outer: it,
		fn:    fn,
	}
}

// Next returns the next flattened element, or (zero, false) when exhausted.
func (f *FlatMapIterator[A, B]) Next() (B, bool) {
	for {
		if f.idx < len(f.inner) {
			val := f.inner[f.idx]
			f.idx++
			return val, true
		}

		outerVal, ok := f.outer.Next()
		if !ok {
			var zero B
			return zero, false
		}

		f.inner = f.fn(outerVal)
		f.idx = 0
	}
}
