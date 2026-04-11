package ops

import "github.com/ver1619/itFrame/core"

// PeekIterator supports lookahead by buffering one element.
type PeekIterator[T any] struct {
	it     core.Iterator[T]
	buf    T
	hasBuf bool
}

// Peek creates a PeekIterator from the given iterator.
func Peek[T any](it core.Iterator[T]) *PeekIterator[T] {
	return &PeekIterator[T]{it: it}
}

// Peek returns the next value without consuming it. Repeated calls return the same value.
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

// Next returns the next element, consuming the peeked value if present.
func (p *PeekIterator[T]) Next() (T, bool) {
	if p.hasBuf {
		p.hasBuf = false
		return p.buf, true
	}
	return p.it.Next()
}
