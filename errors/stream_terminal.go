package errors

func (s Stream[T]) Collect() ([]T, error) {
	return Collect(s.it)
}

/*
- collects values
- stops on first error
*/
