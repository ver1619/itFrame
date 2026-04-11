package ops

import "github.com/ver1619/itFrame/core"

// ZipWithIterator combines two iterators using a custom function, without creating intermediate pairs.
type ZipWithIterator[A, B, C any] struct {
	it1 core.Iterator[A]
	it2 core.Iterator[B]
	fn  func(A, B) C
}

// ZipWith creates an iterator that combines elements from two iterators using fn.
// Stops when either iterator is exhausted.
func ZipWith[A, B, C any](
	it1 core.Iterator[A],
	it2 core.Iterator[B],
	fn func(A, B) C,
) core.Iterator[C] {
	return &ZipWithIterator[A, B, C]{it1: it1, it2: it2, fn: fn}
}

// Next returns the next combined value, or (zero, false) when either iterator is exhausted.
func (z *ZipWithIterator[A, B, C]) Next() (C, bool) {
	v1, ok1 := z.it1.Next()
	v2, ok2 := z.it2.Next()

	if !ok1 || !ok2 {
		var zero C
		return zero, false
	}

	return z.fn(v1, v2), true
}
