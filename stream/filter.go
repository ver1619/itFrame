package stream

import "github.com/ver1619/itFrame/ops"

// Filter returns a new Stream with only elements satisfying pred.
func (s Stream[T]) Filter(pred func(T) bool) Stream[T] {
	return Stream[T]{it: ops.Filter(s.it, pred)}
}
