package core

// SliceIterator iterates over elements of a slice without copying.
type SliceIterator[T any] struct {
	data []T
	idx  int
}

// Slice creates a SliceIterator from the given slice.
func Slice[T any](data []T) *SliceIterator[T] {
	return &SliceIterator[T]{data: data}
}

// Next returns the next element from the slice, or (zero, false) when exhausted.
func (s *SliceIterator[T]) Next() (T, bool) {
	if s.idx >= len(s.data) {
		var zero T
		return zero, false
	}
	val := s.data[s.idx]
	s.idx++
	return val, true
}

// Len returns the number of remaining elements.
func (s *SliceIterator[T]) Len() int {
	remaining := len(s.data) - s.idx
	if remaining < 0 {
		return 0
	}
	return remaining
}
