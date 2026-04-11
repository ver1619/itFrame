package ops

import (
	"github.com/ver1619/itFrame/compare"
	"github.com/ver1619/itFrame/core"
)

// DistinctIterator removes consecutive duplicates from a sorted iterator using a comparator.
type DistinctIterator[T any] struct {
	it      core.Iterator[T]
	cmp     compare.Comparator[T]
	last    T
	hasLast bool
}

// Distinct creates an iterator that removes consecutive duplicates from a sorted iterator.
// Equality is determined by the comparator: !Less(a,b) && !Less(b,a).
func Distinct[T any](it core.Iterator[T], cmp compare.Comparator[T]) core.Iterator[T] {
	return &DistinctIterator[T]{it: it, cmp: cmp}
}

// Next returns the next distinct element, or (zero, false) when exhausted.
func (d *DistinctIterator[T]) Next() (T, bool) {
	for {
		val, ok := d.it.Next()
		if !ok {
			var zero T
			return zero, false
		}

		if d.hasLast && compare.Equal(d.cmp, val, d.last) {
			continue
		}

		d.last = val
		d.hasLast = true
		return val, true
	}
}
