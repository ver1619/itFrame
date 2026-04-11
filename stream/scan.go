package stream

import "github.com/ver1619/itFrame/ops"

// Scan returns a new Stream of running accumulations.
// Each output element is the accumulator after incorporating the corresponding input element.
func (s Stream[T]) Scan(init T, fn func(T, T) T) Stream[T] {
	return Stream[T]{it: ops.Scan(s.it, init, fn)}
}
