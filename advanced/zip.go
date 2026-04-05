package advanced

import "github.com/ver1619/itFrame/core"

type Pair[A, B any] struct {
	First  A
	Second B
}

type ZipIterator[A, B any] struct {
	it1 core.Iterator[A]
	it2 core.Iterator[B]
}

func Zip[A, B any](it1 core.Iterator[A], it2 core.Iterator[B]) *ZipIterator[A, B] {
	return &ZipIterator[A, B]{
		it1: it1,
		it2: it2,
	}
}

func (z *ZipIterator[A, B]) Next() (Pair[A, B], bool) {
	v1, ok1 := z.it1.Next()
	v2, ok2 := z.it2.Next()

	if !ok1 || !ok2 {
		var zero Pair[A, B]
		return zero, false
	}

	return Pair[A, B]{First: v1, Second: v2}, true
}

/*
- **ZipIterator** combines two iterators element-by-element.
- Each call to Next() returns a pair of values.
- Stops when any one iterator is exhausted.
*/
