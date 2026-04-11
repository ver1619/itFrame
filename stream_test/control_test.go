package stream_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/stream"
)

func TestStream_Take(t *testing.T) {
	result := stream.Slice([]int{1, 2, 3, 4, 5}).Take(3).Collect()
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestStream_Skip(t *testing.T) {
	result := stream.Slice([]int{1, 2, 3, 4, 5}).Skip(2).Collect()
	expected := []int{3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestStream_TakeWhile(t *testing.T) {
	result := stream.Slice([]int{2, 4, 6, 7, 8}).
		TakeWhile(func(x int) bool { return x%2 == 0 }).Collect()
	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestStream_DropWhile(t *testing.T) {
	result := stream.Slice([]int{1, 2, 3, 4, 5}).
		DropWhile(func(x int) bool { return x < 3 }).Collect()
	expected := []int{3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestStream_Chain(t *testing.T) {
	s1 := stream.Slice([]int{1, 2})
	s2 := stream.Slice([]int{3, 4})
	result := s1.Chain(s2).Collect()
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestStream_ForEach(t *testing.T) {
	var sum int
	stream.Slice([]int{1, 2, 3}).ForEach(func(x int) { sum += x })
	if sum != 6 {
		t.Fatalf("expected 6, got %d", sum)
	}
}

func TestStream_SkipTake_Composition(t *testing.T) {
	result := stream.Range(0, 100, 1).Skip(10).Take(5).Collect()
	expected := []int{10, 11, 12, 13, 14}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
