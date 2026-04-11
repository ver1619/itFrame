package ops

import "github.com/ver1619/itFrame/core"

// CrossJoinIterator produces the cartesian product of two iterators.
type CrossJoinIterator[A, B any] struct {
	outer core.Iterator[A]
	inner []B

	currA   A
	idx     int
	hasCurr bool
}

// CrossJoin creates a cartesian product iterator. The second iterator is fully buffered.
func CrossJoin[A, B any](
	it1 core.Iterator[A],
	it2 core.Iterator[B],
) core.Iterator[Pair[A, B]] {
	buffer := Collect(it2)

	return &CrossJoinIterator[A, B]{
		outer: it1,
		inner: buffer,
	}
}

// Next returns the next pair in the cartesian product, or (zero, false) when exhausted.
func (c *CrossJoinIterator[A, B]) Next() (Pair[A, B], bool) {
	for {
		if c.hasCurr {
			if c.idx < len(c.inner) {
				p := Pair[A, B]{
					First:  c.currA,
					Second: c.inner[c.idx],
				}
				c.idx++
				return p, true
			}
			c.hasCurr = false
			c.idx = 0
		}

		a, ok := c.outer.Next()
		if !ok {
			var zero Pair[A, B]
			return zero, false
		}

		c.currA = a
		c.hasCurr = true
	}
}
