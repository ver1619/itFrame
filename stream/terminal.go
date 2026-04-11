package stream

import "github.com/ver1619/itFrame/ops"

// Reduce consumes the stream and folds all elements into a single value.
func (s Stream[T]) Reduce(init T, fn func(acc, val T) T) T {
	return ops.Reduce(s.it, init, fn)
}

// Collect consumes the stream and returns all elements as a slice.
func (s Stream[T]) Collect() []T {
	return ops.Collect(s.it)
}

// Count consumes the stream and returns the number of elements.
func (s Stream[T]) Count() int {
	return ops.Count(s.it)
}

// Any returns true if at least one element satisfies pred. Short-circuits on first match.
func (s Stream[T]) Any(pred func(T) bool) bool {
	return ops.Any(s.it, pred)
}

// All returns true if every element satisfies pred. Short-circuits on first failure.
func (s Stream[T]) All(pred func(T) bool) bool {
	return ops.All(s.it, pred)
}
