package stream_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/stream"
)

func TestStream_Map(t *testing.T) {
	s := stream.Slice([]int{1, 2, 3}).
		Map(func(x int) int { return x * 2 })

	result := s.Collect()

	expected := []int{2, 4, 6}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
