package ops

import (
	"github.com/ver1619/itFrame/advanced"
	"github.com/ver1619/itFrame/core"
)

type LeftJoinIterator[A, B, K comparable] struct {
	left core.Iterator[A]
	hash map[K][]B

	leftKey func(A) K

	currA   A
	currBs  []B
	idx     int
	hasCurr bool

	emittedEmpty bool
}

func LeftJoin[A, B, K comparable](
	left core.Iterator[A],
	right core.Iterator[B],
	leftKey func(A) K,
	rightKey func(B) K,
) core.Iterator[advanced.Pair[A, B]] {

	hash := make(map[K][]B)

	// build hash from right
	for {
		v, ok := right.Next()
		if !ok {
			break
		}
		k := rightKey(v)
		hash[k] = append(hash[k], v)
	}

	return &LeftJoinIterator[A, B, K]{
		left:    left,
		hash:    hash,
		leftKey: leftKey,
	}
}

func (j *LeftJoinIterator[A, B, K]) Next() (advanced.Pair[A, B], bool) {
	for {
		// emit matched values
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

		// emit empty (no match case)
		if j.emittedEmpty {
			j.emittedEmpty = false
		}

		a, ok := j.left.Next()
		if !ok {
			var zero advanced.Pair[A, B]
			return zero, false
		}

		k := j.leftKey(a)
		matches, exists := j.hash[k]

		if exists {
			j.currA = a
			j.currBs = matches
			j.hasCurr = true
			continue
		}

		// no match → emit (A, zero B)
		var zeroB B
		return advanced.Pair[A, B]{
			First:  a,
			Second: zeroB,
		}, true
	}
}

/*
- **LeftJoin** returns all left elements, matched with right if available.
- Right side is buffered internally (hash-based).
- If no match → emits left with zero value of right.
- Supports multiple matches (one-to-many)
*/
