// Package core defines the Iterator interface and provides source iterators
// for constructing lazy, single-pass data pipelines.
package core

// Iterator is a pull-based, lazy, single-pass generic iterator.
type Iterator[T any] interface {
	// Next returns the next element and true, or a zero value and false when exhausted.
	Next() (T, bool)
}
