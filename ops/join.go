package ops

import (
	"github.com/ver1619/itFrame/advanced"
	"github.com/ver1619/itFrame/core"
)

type JoinIterator[A, B, K comparable] struct {
	left core.Iterator[A]
	hash map[K][]B

	leftKey func(A) K

	currA   A
	currBs  []B
	idx     int
	hasCurr bool
}

func Join[A, B, K comparable](
	left core.Iterator[A],
	right core.Iterator[B],
	leftKey func(A) K,
	rightKey func(B) K,
) core.Iterator[advanced.Pair[A, B]] {

	hash := make(map[K][]B)

	// build hash table from right iterator
	for {
		v, ok := right.Next()
		if !ok {
			break
		}
		k := rightKey(v)
		hash[k] = append(hash[k], v)
	}

	return &JoinIterator[A, B, K]{
		left:    left,
		hash:    hash,
		leftKey: leftKey,
	}
}

func (j *JoinIterator[A, B, K]) Next() (advanced.Pair[A, B], bool) {
	for {
		if j.hasCurr {
			if j.idx < len(j.currBs) {
				p := advanced.Pair[A, B]{
					First:  j.currA,
					Second: j.currBs[j.idx],
				}
				j.idx++
				return p, true
			}
			j.hasCurr = false
			j.idx = 0
		}

		a, ok := j.left.Next()
		if !ok {
			var zero advanced.Pair[A, B]
			return zero, false
		}

		k := j.leftKey(a)
		matches, exists := j.hash[k]
		if !exists {
			continue
		}

		j.currA = a
		j.currBs = matches
		j.hasCurr = true
	}
}

/*
- **Join** matches elements from two iterators using keys.
- Right iterator is fully buffered internally.
- Left iterator is streamed lazily.
- Only matching pairs are returned (inner join).
*/
