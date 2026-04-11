package stream

import (
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

// MapTo applies a type-changing transformation to each element of the stream.
// This is a free function because Go methods cannot introduce new type parameters.
func MapTo[A, B any](s Stream[A], fn func(A) B) Stream[B] {
	return Stream[B]{it: ops.Map(s.Iterator(), fn)}
}

// Iterator returns the underlying iterator of the stream.
func (s Stream[T]) Iterator() core.Iterator[T] {
	return s.it
}
