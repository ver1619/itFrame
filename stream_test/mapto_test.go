package stream_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/stream"
)

func TestMapTo_IntToString(t *testing.T) {
	s := stream.Slice([]int{1, 2, 3})
	result := stream.MapTo(s, func(x int) string {
		return fmt.Sprintf("v%d", x)
	}).Collect()

	expected := []string{"v1", "v2", "v3"}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestMapTo_StringToInt(t *testing.T) {
	s := stream.Slice([]string{"a", "bb", "ccc"})
	result := stream.MapTo(s, func(x string) int {
		return len(x)
	}).Collect()

	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestMapTo_Chained(t *testing.T) {
	s := stream.Slice([]int{1, 2, 3, 4}).
		Filter(func(x int) bool { return x%2 == 0 })

	result := stream.MapTo(s, func(x int) string {
		return fmt.Sprintf("%d", x)
	}).Collect()

	expected := []string{"2", "4"}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestMapTo_Empty(t *testing.T) {
	s := stream.Slice([]int{})
	result := stream.MapTo(s, func(x int) string {
		return fmt.Sprint(x)
	}).Collect()

	if len(result) != 0 {
		t.Fatalf("expected empty slice, got %v", result)
	}
}

func TestFlatMapTo_IntToStrings(t *testing.T) {
	s := stream.Slice([]int{1, 2, 3})
	result := stream.FlatMapTo(s, func(x int) []string {
		strs := make([]string, x)
		for i := range strs {
			strs[i] = fmt.Sprintf("%d", x)
		}
		return strs
	}).Collect()

	expected := []string{"1", "2", "2", "3", "3", "3"}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestFlatMap_SameType(t *testing.T) {
	s := stream.Slice([]int{1, 2, 3})
	result := s.FlatMap(func(x int) []int {
		return []int{x, x * 10}
	}).Collect()

	expected := []int{1, 10, 2, 20, 3, 30}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestFlatMapTo_Empty(t *testing.T) {
	s := stream.Slice([]int{})
	result := stream.FlatMapTo(s, func(x int) []string {
		return []string{fmt.Sprint(x)}
	}).Collect()

	if len(result) != 0 {
		t.Fatalf("expected empty slice, got %v", result)
	}
}
