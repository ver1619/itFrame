package ops

import "github.com/ver1619/itFrame/core"

// Reduce consumes an iterator and folds all elements into a single result using fn.
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
