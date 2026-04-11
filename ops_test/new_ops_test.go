package ops_test

import (
	"reflect"
	"testing"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

// --- Scan ---

func TestScan_RunningSum(t *testing.T) {
	it := ops.Scan(core.Slice([]int{1, 2, 3, 4}), 0, func(acc, x int) int { return acc + x })
	result := ops.Collect(it)
	expected := []int{1, 3, 6, 10}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestScan_Empty(t *testing.T) {
	it := ops.Scan(core.Slice([]int{}), 0, func(acc, x int) int { return acc + x })
	result := ops.Collect(it)
	if len(result) != 0 {
		t.Fatalf("expected empty, got %v", result)
	}
}

func TestScan_SingleElement(t *testing.T) {
	it := ops.Scan(core.Slice([]int{5}), 10, func(acc, x int) int { return acc + x })
	result := ops.Collect(it)
	expected := []int{15}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

// --- GroupBySorted ---

func TestGroupBySorted_Basic(t *testing.T) {
	it := ops.GroupBySorted(
		core.Slice([]int{1, 1, 2, 2, 2, 3}),
		func(x int) int { return x },
	)

	var keys []int
	var sizes []int
	for {
		g, ok := it.Next()
		if !ok {
			break
		}
		keys = append(keys, g.Key)
		sizes = append(sizes, len(g.Items))
	}

	expectedKeys := []int{1, 2, 3}
	expectedSizes := []int{2, 3, 1}
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Fatalf("expected keys %v, got %v", expectedKeys, keys)
	}
	if !reflect.DeepEqual(sizes, expectedSizes) {
		t.Fatalf("expected sizes %v, got %v", expectedSizes, sizes)
	}
}

func TestGroupBySorted_SingleGroup(t *testing.T) {
	it := ops.GroupBySorted(
		core.Slice([]int{5, 5, 5}),
		func(x int) int { return x },
	)

	g, ok := it.Next()
	if !ok {
		t.Fatal("expected a group")
	}
	if g.Key != 5 || len(g.Items) != 3 {
		t.Fatalf("expected key=5 items=3, got key=%d items=%d", g.Key, len(g.Items))
	}

	_, ok = it.Next()
	if ok {
		t.Fatal("expected exhausted")
	}
}

func TestGroupBySorted_Empty(t *testing.T) {
	it := ops.GroupBySorted(
		core.Slice([]int{}),
		func(x int) int { return x },
	)

	_, ok := it.Next()
	if ok {
		t.Fatal("expected exhausted on empty input")
	}
}

// --- SizedIterator ---

func TestSizedIterator_Slice(t *testing.T) {
	it := core.Slice([]int{1, 2, 3, 4, 5})

	if it.Len() != 5 {
		t.Fatalf("expected Len()=5, got %d", it.Len())
	}

	it.Next()
	it.Next()

	if it.Len() != 3 {
		t.Fatalf("expected Len()=3 after 2 calls, got %d", it.Len())
	}
}

func TestSizedIterator_Range(t *testing.T) {
	it := core.Range(0, 10, 2)

	if it.Len() != 5 {
		t.Fatalf("expected Len()=5, got %d", it.Len())
	}

	it.Next()

	if it.Len() != 4 {
		t.Fatalf("expected Len()=4 after 1 call, got %d", it.Len())
	}
}

func TestSizedIterator_RangeBackward(t *testing.T) {
	it := core.Range(10, 0, -3)
	// values: 10, 7, 4, 1 → 4 elements
	if it.Len() != 4 {
		t.Fatalf("expected Len()=4, got %d", it.Len())
	}
}

func TestCollect_PreAllocates(t *testing.T) {
	it := core.Slice([]int{1, 2, 3})
	result := ops.Collect(it)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
