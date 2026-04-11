// Package errors provides error-aware iteration using Result[T] values,
// enabling safe data processing pipelines with automatic error propagation.
package errors

// Result wraps a value or an error for use in error-aware pipelines.
type Result[T any] struct {
	Value T
	Err   error
}

// Ok creates a successful Result containing v.
func Ok[T any](v T) Result[T] {
	return Result[T]{Value: v}
}

// ErrResult creates an error Result containing err.
func ErrResult[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

// IsError returns true if the Result contains an error.
func (r Result[T]) IsError() bool {
	return r.Err != nil
}
