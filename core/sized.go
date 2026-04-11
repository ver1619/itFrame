package core

// SizedIterator is an optional interface for iterators that know their remaining length.
// Implementations of Collect use this to pre-allocate slices.
type SizedIterator interface {
	Len() int
}
