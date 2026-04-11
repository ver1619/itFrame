// Package stream provides a fluent, chainable API over iterators for composable data pipelines.
package stream

import "github.com/ver1619/itFrame/core"

// Stream wraps an iterator and enables method chaining for lazy pipelines.
type Stream[T any] struct {
	it core.Iterator[T]
}

// From creates a Stream from any iterator.
func From[T any](it core.Iterator[T]) Stream[T] {
	return Stream[T]{it: it}
}

// Slice creates a Stream from a slice.
func Slice[T any](data []T) Stream[T] {
	return Stream[T]{it: core.Slice(data)}
}

// Range creates a Stream of integers over [start, end) with the given step.
func Range(start, end, step int) Stream[int] {
	return Stream[int]{it: core.Range(start, end, step)}
}
