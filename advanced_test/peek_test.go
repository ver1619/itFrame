package advanced_test

import (
	"testing"

	"github.com/ver1619/itFrame/advanced"
	"github.com/ver1619/itFrame/core"
)

func TestPeek_Basic(t *testing.T) {
	p := advanced.Peek(core.Slice([]int{1, 2, 3}))

	v, ok := p.Peek()
	if !ok || v != 1 {
		t.Fatalf("expected 1, got %v", v)
	}

	v, ok = p.Next()
	if !ok || v != 1 {
		t.Fatalf("expected 1 after peek, got %v", v)
	}
}

func TestPeek_MultiplePeek(t *testing.T) {
	p := advanced.Peek(core.Slice([]int{10, 20}))

	v1, _ := p.Peek()
	v2, _ := p.Peek()

	if v1 != v2 {
		t.Fatal("peek should return same value repeatedly")
	}
}

func TestPeek_PeekThenNextSequence(t *testing.T) {
	p := advanced.Peek(core.Slice([]int{1, 2}))

	v, _ := p.Peek()
	if v != 1 {
		t.Fatal("expected 1")
	}

	v, _ = p.Next()
	if v != 1 {
		t.Fatal("expected 1 again")
	}

	v, _ = p.Next()
	if v != 2 {
		t.Fatal("expected 2")
	}
}

func TestPeek_Exhaustion(t *testing.T) {
	p := advanced.Peek(core.Slice([]int{1}))

	p.Next()

	_, ok := p.Peek()
	if ok {
		t.Fatal("expected exhaustion")
	}

	_, ok = p.Next()
	if ok {
		t.Fatal("expected exhaustion")
	}
}
