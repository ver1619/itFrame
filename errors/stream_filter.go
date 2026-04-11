package errors

// Filter returns a new error-aware Stream with only valid values satisfying pred.
// Errors are preserved and passed through unchanged.
func (s Stream[T]) Filter(pred func(T) bool) Stream[T] {
	it := Filter[T](s.it, pred)
	return Stream[T]{it: it}
}
