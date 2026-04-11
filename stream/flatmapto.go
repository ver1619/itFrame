package stream

import (
	"github.com/ver1619/itFrame/ops"
)

// FlatMapTo is a free function that flat-maps a stream of type A to a stream of type B returning slices.
// Because it works with slices instead of iterators, it is significantly more memory-efficient.
func FlatMapTo[A, B any](s Stream[A], fn func(A) []B) Stream[B] {
	return Stream[B]{it: ops.FlatMap(s.Iterator(), fn)}
}
