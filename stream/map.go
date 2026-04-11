package stream

import "github.com/ver1619/itFrame/ops"

// Map returns a new Stream with each element transformed by fn.
// For type-changing transformations (A → B), use the free function MapTo.
func (s Stream[T]) Map(fn func(T) T) Stream[T] {
	return Stream[T]{it: ops.Map(s.it, fn)}
}
