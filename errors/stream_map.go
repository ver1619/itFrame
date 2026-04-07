package errors

func (s Stream[T]) Map(fn func(T) T) Stream[T] {
	it := Map(s.it, fn)
	return Stream[T]{it: it}
}

/*
- applies function only on valid values
- errors pass through unchanged
*/
