package compare

// Equal returns true if a and b are equivalent under the given comparator.
func Equal[T any](cmp Comparator[T], a, b T) bool {
	return !cmp.Less(a, b) && !cmp.Less(b, a)
}

// LessOrEqual returns true if a is less than or equal to b under the given comparator.
func LessOrEqual[T any](cmp Comparator[T], a, b T) bool {
	return !cmp.Less(b, a)
}
