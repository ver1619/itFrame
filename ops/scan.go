package ops

import "github.com/ver1619/itFrame/core"

// ScanIterator emits running accumulations of the input elements.
type ScanIterator[T any, R any] struct {
	it  core.Iterator[T]
	acc R
	fn  func(R, T) R
}

// Scan creates a ScanIterator that emits the accumulator value after incorporating each element.
// Unlike Reduce, Scan is lazy and yields one output per input element.
func Scan[T any, R any](it core.Iterator[T], init R, fn func(R, T) R) core.Iterator[R] {
	return &ScanIterator[T, R]{it: it, acc: init, fn: fn}
}

// Next returns the next accumulated value, or (zero, false) when exhausted.
func (s *ScanIterator[T, R]) Next() (R, bool) {
	val, ok := s.it.Next()
	if !ok {
		var zero R
		return zero, false
	}
	s.acc = s.fn(s.acc, val)
	return s.acc, true
}
