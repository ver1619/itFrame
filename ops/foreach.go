package ops

import "github.com/ver1619/itFrame/core"

// ForEach consumes an iterator and applies fn to each element.
func ForEach[T any](it core.Iterator[T], fn func(T)) {
	for {
		val, ok := it.Next()
		if !ok {
			return
		}
		fn(val)
	}
}
