package advanced

import "github.com/ver1619/itFrame/core"

type ZipLongestIterator[A, B any] struct {
	it1 core.Iterator[A]
	it2 core.Iterator[B]
}

func ZipLongest[A, B any](
	it1 core.Iterator[A],
	it2 core.Iterator[B],
) *ZipLongestIterator[A, B] {
	return &ZipLongestIterator[A, B]{
		it1: it1,
		it2: it2,
	}
}

func (z *ZipLongestIterator[A, B]) Next() (Pair[A, B], bool) {
	v1, ok1 := z.it1.Next()
	v2, ok2 := z.it2.Next()

	if !ok1 && !ok2 {
		var zero Pair[A, B]
		return zero, false
	}

	return Pair[A, B]{First: v1, Second: v2}, true
}
