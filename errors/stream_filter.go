package errors

func (s Stream[T]) Filter(pred func(T) bool) Stream[T] {
	it := Filter[T](s.it, pred)
	return Stream[T]{it: it}

}

/*
- filters valid values
- errors are preserved
*/
