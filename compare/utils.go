package compare

func Equal[T any](cmp Comparator[T], a, b T) bool {
	return !cmp.Less(a, b) && !cmp.Less(b, a)
}

func LessOrEqual[T any](cmp Comparator[T], a, b T) bool {
	return !cmp.Less(b, a)
}

// **Equal** derives equality from ordering
