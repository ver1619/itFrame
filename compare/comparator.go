// Package compare provides a Comparator interface and utilities for ordering-aware operations.
package compare

// Comparator defines a strict weak ordering via Less.
type Comparator[T any] interface {
	Less(a, b T) bool
}

// LessFunc is an adapter that allows using a function as a Comparator.
type LessFunc[T any] func(a, b T) bool

// Less calls the underlying function.
func (f LessFunc[T]) Less(a, b T) bool {
	return f(a, b)
}
