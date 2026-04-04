package stream

import "github.com/ver1619/itFrame/ops"

func (s Stream[T]) Reduce(init T, fn func(acc, val T) T) T {
	return ops.Reduce(s.it, init, fn)
}

func (s Stream[T]) Collect() []T {
	return ops.Collect(s.it)
}

func (s Stream[T]) Count() int {
	return ops.Count(s.it)
}

func (s Stream[T]) Any(pred func(T) bool) bool {
	return ops.Any(s.it, pred)
}

func (s Stream[T]) All(pred func(T) bool) bool {
	return ops.All(s.it, pred)
}
