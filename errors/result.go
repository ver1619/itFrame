package errors

type Result[T any] struct {
	Value T
	Err   error
}

func Ok[T any](v T) Result[T] {
	return Result[T]{Value: v}
}

func ErrResult[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

func (r Result[T]) IsError() bool {
	return r.Err != nil
}

/*
- **Result[T]** wraps a value or an error
- **Ok(v)** → successful value
- **ErrResult(err)** → error case
- **IsError()** checks if error exists
*/
