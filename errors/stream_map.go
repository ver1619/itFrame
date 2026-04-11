package errors

// Map returns a new error-aware Stream with valid values transformed by fn.
// Errors pass through unchanged.
func (s Stream[T]) Map(fn func(T) T) Stream[T] {
	it := Map(s.it, fn)
	return Stream[T]{it: it}
}
