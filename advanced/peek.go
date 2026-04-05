package advanced

import "github.com/ver1619/itFrame/core"

type PeekIterator[T any] struct {
	it     core.Iterator[T]
	buf    T
	hasBuf bool
}

func Peek[T any](it core.Iterator[T]) *PeekIterator[T] {
	return &PeekIterator[T]{it: it}
}

func (p *PeekIterator[T]) Peek() (T, bool) {
	if p.hasBuf {
		return p.buf, true
	}

	val, ok := p.it.Next()
	if !ok {
		var zero T
		return zero, false
	}

	p.buf = val
	p.hasBuf = true
	return val, true
}

func (p *PeekIterator[T]) Next() (T, bool) {
	if p.hasBuf {
		p.hasBuf = false
		return p.buf, true
	}
	return p.it.Next()
}

/*

- **PeekIterator** lets you look at the next value without consuming it.
- **Peek()** returns the next element but does not move the iterator.
- **Next()** behaves normally, but returns the peeked value if available.
- Works lazily — no preloading of data.

*/
