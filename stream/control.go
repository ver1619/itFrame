package stream

import "github.com/ver1619/itFrame/ops"

// Take returns a new Stream that yields at most n elements.
func (s Stream[T]) Take(n int) Stream[T] {
	return Stream[T]{it: ops.Take(s.it, n)}
}

// Skip returns a new Stream that discards the first n elements.
func (s Stream[T]) Skip(n int) Stream[T] {
	return Stream[T]{it: ops.Skip(s.it, n)}
}

// TakeWhile returns a new Stream that yields elements while pred is satisfied.
func (s Stream[T]) TakeWhile(pred func(T) bool) Stream[T] {
	return Stream[T]{it: ops.TakeWhile(s.it, pred)}
}

// DropWhile returns a new Stream that skips elements while pred is satisfied, then yields the rest.
func (s Stream[T]) DropWhile(pred func(T) bool) Stream[T] {
	return Stream[T]{it: ops.DropWhile(s.it, pred)}
}

// Chain returns a new Stream that yields all elements from this stream, then all from other.
func (s Stream[T]) Chain(other Stream[T]) Stream[T] {
	return Stream[T]{it: ops.Chain(s.it, other.it)}
}

// ForEach consumes the stream and applies fn to each element.
func (s Stream[T]) ForEach(fn func(T)) {
	ops.ForEach(s.it, fn)
}
