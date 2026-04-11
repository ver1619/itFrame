package ops

import "github.com/ver1619/itFrame/core"

// Collect consumes an iterator and returns all elements as a slice.
// If the iterator implements core.SizedIterator, the slice is pre-allocated.
func Collect[T any](it core.Iterator[T]) []T {
	var result []T

	if sized, ok := it.(core.SizedIterator); ok {
		result = make([]T, 0, sized.Len())
	}

	for {
		val, ok := it.Next()
		if !ok {
			break
		}
		result = append(result, val)
	}

	return result
}
