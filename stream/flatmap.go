package stream

import "github.com/ver1619/itFrame/ops"

// FlatMap maps each element to a slice and flattens the results into a new Stream.
// This is typically more performant than FlatMap because it avoids allocating new iterators per mapping.
func (s Stream[T]) FlatMap(fn func(T) []T) Stream[T] {
	return Stream[T]{it: ops.FlatMap(s.it, fn)}
}
