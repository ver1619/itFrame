package compare

type Reverse[T any] struct {
	Cmp Comparator[T]
}

func (r Reverse[T]) Less(a, b T) bool {
	return r.Cmp.Less(b, a)
}

// **Reverse** flips ordering
