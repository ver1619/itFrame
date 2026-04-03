package ops

import "github.com/ver1619/itFrame/core"

// terminal operation
func Reduce[T any, R any](it core.Iterator[T], init R, fn func(R, T) R) R {
	acc := init

	for {
		val, ok := it.Next()
		if !ok {
			break
		}
		acc = fn(acc, val)
	}

	return acc
}

/*
- **Reduce** consumes an iterator and produces a single result.
- init is the starting value of the accumulator.
- fn defines how each element updates the result.
- Iteration continues until exhaustion.
- Returns the final accumulated value.
- Does not allocate memory during iteration.
- The iterator is fully consumed after execution.
*/
