package compare

type Multi[T any] struct {
	Comparators []Comparator[T]
}

func (m Multi[T]) Less(a, b T) bool {
	for _, cmp := range m.Comparators {
		if cmp.Less(a, b) {
			return true
		}
		if cmp.Less(b, a) {
			return false
		}
	}
	return false
}

// **Multi** supports multi-field sorting
