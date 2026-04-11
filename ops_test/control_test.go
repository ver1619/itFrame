package ops_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func TestTake_Basic(t *testing.T) {
	result := ops.Collect(ops.Take(core.Slice([]int{1, 2, 3, 4, 5}), 3))
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestTake_MoreThanAvailable(t *testing.T) {
	result := ops.Collect(ops.Take(core.Slice([]int{1, 2}), 5))
	expected := []int{1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestTake_Zero(t *testing.T) {
	result := ops.Collect(ops.Take(core.Slice([]int{1, 2, 3}), 0))
	if len(result) != 0 {
		t.Fatalf("expected empty, got %v", result)
	}
}

func TestSkip_Basic(t *testing.T) {
	result := ops.Collect(ops.Skip(core.Slice([]int{1, 2, 3, 4, 5}), 2))
	expected := []int{3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestSkip_MoreThanAvailable(t *testing.T) {
	result := ops.Collect(ops.Skip(core.Slice([]int{1, 2}), 5))
	if len(result) != 0 {
		t.Fatalf("expected empty, got %v", result)
	}
}

func TestSkip_Zero(t *testing.T) {
	result := ops.Collect(ops.Skip(core.Slice([]int{1, 2, 3}), 0))
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestTakeWhile_Basic(t *testing.T) {
	result := ops.Collect(ops.TakeWhile(
		core.Slice([]int{1, 2, 3, 4, 5}),
		func(x int) bool { return x < 4 },
	))
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestTakeWhile_AllMatch(t *testing.T) {
	result := ops.Collect(ops.TakeWhile(
		core.Slice([]int{1, 2, 3}),
		func(x int) bool { return true },
	))
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestTakeWhile_NoneMatch(t *testing.T) {
	result := ops.Collect(ops.TakeWhile(
		core.Slice([]int{1, 2, 3}),
		func(x int) bool { return false },
	))
	if len(result) != 0 {
		t.Fatalf("expected empty, got %v", result)
	}
}

func TestDropWhile_Basic(t *testing.T) {
	result := ops.Collect(ops.DropWhile(
		core.Slice([]int{1, 2, 3, 4, 5}),
		func(x int) bool { return x < 3 },
	))
	expected := []int{3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestDropWhile_AllMatch(t *testing.T) {
	result := ops.Collect(ops.DropWhile(
		core.Slice([]int{1, 2, 3}),
		func(x int) bool { return true },
	))
	if len(result) != 0 {
		t.Fatalf("expected empty, got %v", result)
	}
}

func TestDropWhile_NoneMatch(t *testing.T) {
	result := ops.Collect(ops.DropWhile(
		core.Slice([]int{1, 2, 3}),
		func(x int) bool { return false },
	))
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestChain_Basic(t *testing.T) {
	result := ops.Collect(ops.Chain(
		core.Slice([]int{1, 2}),
		core.Slice([]int{3, 4}),
	))
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestChain_FirstEmpty(t *testing.T) {
	result := ops.Collect(ops.Chain(
		core.Slice([]int{}),
		core.Slice([]int{1, 2}),
	))
	expected := []int{1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestChain_SecondEmpty(t *testing.T) {
	result := ops.Collect(ops.Chain(
		core.Slice([]int{1, 2}),
		core.Slice([]int{}),
	))
	expected := []int{1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestChain_BothEmpty(t *testing.T) {
	result := ops.Collect(ops.Chain(
		core.Slice([]int{}),
		core.Slice([]int{}),
	))
	if len(result) != 0 {
		t.Fatalf("expected empty, got %v", result)
	}
}

func TestForEach_Basic(t *testing.T) {
	var collected []int
	ops.ForEach(core.Slice([]int{1, 2, 3}), func(x int) {
		collected = append(collected, x)
	})
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(collected, expected) {
		t.Fatalf("expected %v, got %v", expected, collected)
	}
}

func TestForEach_Empty(t *testing.T) {
	called := false
	ops.ForEach(core.Slice([]int{}), func(x int) {
		called = true
	})
	if called {
		t.Fatal("expected fn not to be called on empty iterator")
	}
}
