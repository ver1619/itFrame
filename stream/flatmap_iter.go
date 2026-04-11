package stream

import (
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

// FlatMapIter maps each element to an iterator and flattens the results into a new Stream.
// For type-changing flat-map (A → Iterator[B]), use the free function FlatMapIterTo.
func (s Stream[T]) FlatMapIter(fn func(T) core.Iterator[T]) Stream[T] {
	return Stream[T]{it: ops.FlatMapIter(s.it, fn)}
}
