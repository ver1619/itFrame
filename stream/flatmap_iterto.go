package stream

import (
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

// FlatMapIterTo applies a type-changing flat-map over the stream.
// This is a free function because Go methods cannot introduce new type parameters.
func FlatMapIterTo[A, B any](s Stream[A], fn func(A) core.Iterator[B]) Stream[B] {
	return Stream[B]{it: ops.FlatMapIter(s.Iterator(), fn)}
}
