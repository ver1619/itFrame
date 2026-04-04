package stream_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/stream"
)

func TestStream_Composition(t *testing.T) {
	result := stream.Slice([]int{1, 2, 3, 4}).
		Map(func(x int) int { return x * x }).
		Filter(func(x int) bool { return x%2 == 0 }).
		Collect()

	expected := []int{4, 16}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
