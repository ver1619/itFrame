package ops

import "github.com/ver1619/itFrame/core"

// terminal operation
func Collect[T any](it core.Iterator[T]) []T {
	var result []T

	for {
		val, ok := it.Next()
		if !ok {
			break
		}
		result = append(result, val)
	}

	return result
}

/*
- **Collect** converts an iterator into a slice.
- It consumes the iterator fully.
- Each element is appended to a resulting slice.
- Returns all values in the form of a slice and in order.
- Works with any iterator (Slice, Range, Map, Filter, etc.).
- **Collect** allocates memory to store elements as slice grows dynamically using append
*/
