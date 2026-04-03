package ops

import "github.com/ver1619/itFrame/core"

// terminal operation
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

/*
**Any**
- Returns true if at least one element satisfies the condition.
- Stops immediately when a match is found (short-circuit).
- If no elements match → returns false.

**All**
- Returns true if all elements satisfy the condition.
- Stops immediately when a failure is found.
- For empty iterators → returns true.
*/
