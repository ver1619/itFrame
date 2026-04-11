package ops

import "github.com/ver1619/itFrame/core"

// ZipIterator pairs elements from two iterators element-by-element.
// Stops when either iterator is exhausted.
type ZipIterator[A, B any] struct {
	it1 core.Iterator[A]
	it2 core.Iterator[B]
}

// Zip creates an iterator that combines two iterators into pairs.
func Zip[A, B any](it1 core.Iterator[A], it2 core.Iterator[B]) core.Iterator[Pair[A, B]] {
	return &ZipIterator[A, B]{it1: it1, it2: it2}
}

// Next returns the next pair, or (zero, false) when either iterator is exhausted.
func (z *ZipIterator[A, B]) Next() (Pair[A, B], bool) {
	v1, ok1 := z.it1.Next()
	v2, ok2 := z.it2.Next()

	if !ok1 || !ok2 {
		var zero Pair[A, B]
		return zero, false
	}

	return Pair[A, B]{First: v1, Second: v2}, true
}
