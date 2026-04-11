package errors

// Collect consumes the error-aware stream and returns all valid values as a slice.
// Stops immediately on the first error and returns it.
func (s Stream[T]) Collect() ([]T, error) {
	return Collect(s.it)
}
