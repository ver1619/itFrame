package stream

import "github.com/ver1619/itFrame/ops"

func (s Stream[T]) Filter(pred func(T) bool) Stream[T] {
	return Stream[T]{it: ops.Filter(s.it, pred)}
}

/*
**Filter**
- Keeps only elements that satisfy a condition.
- Returns a new stream.
- Evaluation happens only when a terminal operation is called.
*/
