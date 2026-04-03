package ops

import "github.com/ver1619/itFrame/core"

// terminal operation
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

/*
- **Count** returns the number of elements in an iterator.
- It consumes the iterator completely.
- Each call to Next() increments a counter.
- Returns total number of elements.
- Does not allocate memory.
*/
