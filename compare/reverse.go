package compare

// Reverse wraps a Comparator and inverts its ordering.
type Reverse[T any] struct {
	Cmp Comparator[T]
}

// Less returns true if b is less than a according to the wrapped comparator.
func (r Reverse[T]) Less(a, b T) bool {
	return r.Cmp.Less(b, a)
}
