package advanced

import "github.com/ver1619/itFrame/core"

type ZipWithIterator[A, B, C any] struct {
	it1 core.Iterator[A]
	it2 core.Iterator[B]
	fn  func(A, B) C
}

func ZipWith[A, B, C any](
	it1 core.Iterator[A],
	it2 core.Iterator[B],
	fn func(A, B) C,
) *ZipWithIterator[A, B, C] {
	return &ZipWithIterator[A, B, C]{
		it1: it1,
		it2: it2,
		fn:  fn,
	}
}

func (z *ZipWithIterator[A, B, C]) Next() (C, bool) {
	v1, ok1 := z.it1.Next()
	v2, ok2 := z.it2.Next()

	if !ok1 || !ok2 {
		var zero C
		return zero, false
	}

	return z.fn(v1, v2), true
}

/*
- **ZipWith** combines two iterators using a function.
- Each pair of values is transformed into a new value.
- Stops when any iterator is exhausted.
- No intermediate Pair is created.
*/
