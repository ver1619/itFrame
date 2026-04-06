package ops

import (
	"github.com/ver1619/itFrame/advanced"
	"github.com/ver1619/itFrame/core"
)

type CrossJoinIterator[A, B any] struct {
	outer core.Iterator[A]
	inner []B

	currA   A
	idx     int
	hasCurr bool
}

func CrossJoin[A, B any](
	it1 core.Iterator[A],
	it2 core.Iterator[B],
) core.Iterator[advanced.Pair[A, B]] {
	// buffer inner completely
	buffer := Collect(it2)

	return &CrossJoinIterator[A, B]{
		outer: it1,
		inner: buffer,
	}
}

func (c *CrossJoinIterator[A, B]) Next() (advanced.Pair[A, B], bool) {
	for {
		// if we have active outer value
		if c.hasCurr {
			if c.idx < len(c.inner) {
				p := advanced.Pair[A, B]{
					First:  c.currA,
					Second: c.inner[c.idx],
				}
				c.idx++
				return p, true
			}
			// reset for next outer
			c.hasCurr = false
			c.idx = 0
		}

		// fetch next outer
		a, ok := c.outer.Next()
		if !ok {
			var zero advanced.Pair[A, B]
			return zero, false
		}

		c.currA = a
		c.hasCurr = true
	}
}

/*
- **CrossJoin** produces all combinations between two iterators.
- Second iterator is fully buffered internally.
- Output size = len(A) × len(B).
- Lazy over outer, but inner is preloaded.
*/
