package ops

import "github.com/ver1619/itFrame/core"

// LeftJoinIterator performs a left join between two iterators using key functions.
// Unmatched left elements are emitted with a zero-value right element.
type LeftJoinIterator[A, B, K comparable] struct {
	left core.Iterator[A]
	hash map[K][]B

	leftKey func(A) K

	currA   A
	currBs  []B
	idx     int
	hasCurr bool
}

// LeftJoin creates a left join iterator. The right iterator is fully buffered into a hash map.
// All left elements are returned; unmatched left elements are paired with a zero-value right.
func LeftJoin[A, B, K comparable](
	left core.Iterator[A],
	right core.Iterator[B],
	leftKey func(A) K,
	rightKey func(B) K,
) core.Iterator[Pair[A, B]] {

	hash := make(map[K][]B)

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

// Next returns the next pair, or (zero, false) when exhausted.
func (j *LeftJoinIterator[A, B, K]) Next() (Pair[A, B], bool) {
	for {
		if j.hasCurr {
			if j.idx < len(j.currBs) {
				p := Pair[A, B]{
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
			var zero Pair[A, B]
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

		var zeroB B
		return Pair[A, B]{
			First:  a,
			Second: zeroB,
		}, true
	}
}
