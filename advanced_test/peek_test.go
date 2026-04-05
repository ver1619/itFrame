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
		t.Fatalf("expected 1 from Next after Peek")
	}
}

func TestPeek_MultiplePeek(t *testing.T) {
	p := advanced.Peek(core.Slice([]int{1, 2}))

	v1, _ := p.Peek()
	v2, _ := p.Peek()

	if v1 != v2 {
		t.Fatal("peek should return same value")
	}
}

func TestPeek_Exhaustion(t *testing.T) {
	p := advanced.Peek(core.Slice([]int{1}))

	p.Next()

	_, ok := p.Peek()
	if ok {
		t.Fatal("expected exhaustion")
	}
}
