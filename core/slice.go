package core

type SliceIterator[T any] struct {
	data []T
	idx  int
}

func NewSliceIterator[T any](data []T) *SliceIterator[T] {
	return &SliceIterator[T]{data: data}
}

func (s *SliceIterator[T]) Next() (T, bool) {
	if s.idx >= len(s.data) {
		var zero T
		return zero, false
	}
	val := s.data[s.idx]
	s.idx++
	return val, true
}

/*
SliceIterator lets you iterate over a slice one element at a time using Next().
NewSliceIterator(data) creates an iterator without copying the slice.
*/
