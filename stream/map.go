package stream

import "github.com/ver1619/itFrame/ops"

func (s Stream[T]) Map(fn func(T) T) Stream[T] {
	return Stream[T]{it: ops.Map(s.it, fn)}
}

/*
**Map**
- Transforms each element in the stream.
- Returns a new stream with transformed values.
- Does not execute immediately (lazy).
*/
