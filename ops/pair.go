package ops

// Pair is a generic tuple holding two values.
type Pair[A, B any] struct {
	First  A
	Second B
}
