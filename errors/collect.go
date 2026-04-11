package errors

import "github.com/ver1619/itFrame/core"

// Collect consumes an iterator of Result values and returns all valid values as a slice.
// Stops immediately on the first error and returns it.
func Collect[T any](it core.Iterator[Result[T]]) ([]T, error) {
	var result []T

	for {
		r, ok := it.Next()
		if !ok {
			break
		}

		if r.Err != nil {
			return nil, r.Err
		}

		result = append(result, r.Value)
	}

	return result, nil
}
