package ops

import "github.com/ver1619/itFrame/core"

// Any returns true if at least one element satisfies pred. Short-circuits on first match.
func Any[T any](it core.Iterator[T], pred func(T) bool) bool {
	for {
		val, ok := it.Next()
		if !ok {
			return false
		}
		if pred(val) {
			return true
		}
	}
}

// All returns true if every element satisfies pred. Short-circuits on first failure.
func All[T any](it core.Iterator[T], pred func(T) bool) bool {
	for {
		val, ok := it.Next()
		if !ok {
			return true
		}
		if !pred(val) {
			return false
		}
	}
}
