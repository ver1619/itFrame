package stream_test

import (
	"testing"

	"github.com/ver1619/itFrame/stream"
)

func TestStream_Reduce(t *testing.T) {
	result := stream.Slice([]int{1, 2, 3, 4}).
		Reduce(0, func(acc, x int) int { return acc + x })

	if result != 10 {
		t.Fatalf("expected 10, got %d", result)
	}
}

func TestStream_Count(t *testing.T) {
	c := stream.Range(0, 5, 1).Count()

	if c != 5 {
		t.Fatalf("expected 5, got %d", c)
	}
}

func TestStream_Any(t *testing.T) {
	ok := stream.Range(0, 10, 1).
		Any(func(x int) bool { return x == 5 })

	if !ok {
		t.Fatal("expected true")
	}
}

func TestStream_All(t *testing.T) {
	ok := stream.Slice([]int{2, 4, 6}).
		All(func(x int) bool { return x%2 == 0 })

	if !ok {
		t.Fatal("expected true")
	}
}
