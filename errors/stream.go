package errors

import (
	"github.com/ver1619/itFrame/core"
)

// Stream wraps an iterator of Result values for error-aware method chaining.
type Stream[T any] struct {
	it core.Iterator[Result[T]]
}

// FromIterator creates an error-aware Stream from an iterator of Result values.
func FromIterator[T any](it core.Iterator[Result[T]]) Stream[T] {
	return Stream[T]{it: it}
}

// FromSlice creates an error-aware Stream from a slice of Result values.
func FromSlice[T any](data []Result[T]) Stream[T] {
	return Stream[T]{it: core.Slice(data)}
}

// Iterator returns the underlying iterator of the stream.
func (s Stream[T]) Iterator() core.Iterator[Result[T]] {
	return s.it
}
