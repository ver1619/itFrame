package advanced

// Pair is a generic structural type that holds two values.
type Pair[A, B any] struct {
	First  A
	Second B
}
