package errors

import (
	"github.com/ver1619/itFrame/core"
)

type Stream[T any] struct {
	it core.Iterator[Result[T]]
}

func FromIterator[T any](it core.Iterator[Result[T]]) Stream[T] {
	return Stream[T]{it: it}
}

func FromSlice[T any](data []Result[T]) Stream[T] {
	return Stream[T]{it: core.Slice(data)}
}

func (s Stream[T]) Iterator() core.Iterator[Result[T]] {
	return s.it
}

/*
- **errors.Stream[T]** works with **Result[T]**
- wraps iterator of results
- base entry point for error-aware pipelines
*/
