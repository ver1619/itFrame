package stream

import "github.com/ver1619/itFrame/core"

type Stream[T any] struct {
	it core.Iterator[T]
}

func From[T any](it core.Iterator[T]) Stream[T] {
	return Stream[T]{it: it}
}

func Slice[T any](data []T) Stream[T] {
	return Stream[T]{it: core.Slice(data)}
}

func Range(start, end, step int) Stream[int] {
	return Stream[int]{it: core.Range(start, end, step)}
}

/*
- **Stream[T]** is a wrapper over an iterator that enables method chaining.
- From(it) converts any iterator into a stream.
- Slice(data) creates a stream from a slice.
- Range(start, end, step) creates a numeric stream.
No data is processed during creation — execution starts only when terminal methods are called.
*/
