package compare

// Multi combines multiple comparators for multi-field sorting.
// Comparators are evaluated in order; the first decisive result wins.
type Multi[T any] struct {
	Comparators []Comparator[T]
}

// Less returns true if a is less than b according to the first decisive comparator.
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
