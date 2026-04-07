package errors

import "github.com/ver1619/itFrame/core"

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

/*
- collect values until error occurs
- stops immediately on first error
*/
