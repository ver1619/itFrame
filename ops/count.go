package ops

import "github.com/ver1619/itFrame/core"

// Count consumes an iterator and returns the number of elements.
func Count[T any](it core.Iterator[T]) int {
	count := 0

	for {
		_, ok := it.Next()
		if !ok {
			break
		}
		count++
	}

	return count
}
