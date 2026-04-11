package ops

import "github.com/ver1619/itFrame/core"

// ZipLongestIterator pairs elements from two iterators, continuing until both are exhausted.
// Missing values are filled with zero values.
type ZipLongestIterator[A, B any] struct {
	it1 core.Iterator[A]
	it2 core.Iterator[B]
}

// ZipLongest creates an iterator that pairs elements from two iterators.
// Continues until both iterators are exhausted, using zero values for the shorter one.
func ZipLongest[A, B any](
	it1 core.Iterator[A],
	it2 core.Iterator[B],
) core.Iterator[Pair[A, B]] {
	return &ZipLongestIterator[A, B]{it1: it1, it2: it2}
}

// Next returns the next pair, or (zero, false) when both iterators are exhausted.
func (z *ZipLongestIterator[A, B]) Next() (Pair[A, B], bool) {
	v1, ok1 := z.it1.Next()
	v2, ok2 := z.it2.Next()

	if !ok1 && !ok2 {
		var zero Pair[A, B]
		return zero, false
	}

	return Pair[A, B]{First: v1, Second: v2}, true
}
