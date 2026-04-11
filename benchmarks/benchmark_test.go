package benchmarks_test

import (
	"testing"

	"github.com/ver1619/itFrame/compare"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

var sink int
var sinkSlice []int

func generateData(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	return data
}

func BenchmarkMap(b *testing.B) {
	data := generateData(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		it := ops.Map(core.Slice(data), func(x int) int {
			return x * 2
		})

		sum := 0
		for {
			v, ok := it.Next()
			if !ok {
				break
			}
			sum += v
		}
		sink = sum
	}
}

func BenchmarkFilter(b *testing.B) {
	data := generateData(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		it := ops.Filter(core.Slice(data), func(x int) bool {
			return x%2 == 0
		})

		sum := 0
		for {
			v, ok := it.Next()
			if !ok {
				break
			}
			sum += v
		}
		sink = sum
	}
}

func BenchmarkFlatMapIter(b *testing.B) {
	data := generateData(1000)
	inner := []int{1, 2, 3}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		it := ops.FlatMapIter(core.Slice(data), func(x int) core.Iterator[int] {
			return core.Slice(inner)
		})

		sum := 0
		for {
			v, ok := it.Next()
			if !ok {
				break
			}
			sum += v
		}
		sink = sum
	}
}

func BenchmarkFlatMap(b *testing.B) {
	data := generateData(1000)
	inner := []int{1, 2, 3}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		it := ops.FlatMap(core.Slice(data), func(x int) []int {
			return inner
		})

		sum := 0
		for {
			v, ok := it.Next()
			if !ok {
				break
			}
			sum += v
		}
		sink = sum
	}
}

func BenchmarkMerge(b *testing.B) {
	data1 := generateData(5000)
	data2 := generateData(5000)
	cmp := compare.LessFunc[int](func(a, b int) bool { return a < b })
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		it := ops.Merge(core.Slice(data1), core.Slice(data2), cmp)

		sum := 0
		for {
			v, ok := it.Next()
			if !ok {
				break
			}
			sum += v
		}
		sink = sum
	}
}

func BenchmarkJoin(b *testing.B) {
	left := generateData(1000)
	right := generateData(500)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		it := ops.Join(
			core.Slice(left),
			core.Slice(right),
			func(x int) int { return x % 100 },
			func(x int) int { return x % 100 },
		)

		count := 0
		for {
			_, ok := it.Next()
			if !ok {
				break
			}
			count++
		}
		sink = count
	}
}

func BenchmarkCollectSized(b *testing.B) {
	data := generateData(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Slice iterator implements SizedIterator, triggering preallocation
		res := ops.Collect(core.Slice(data))
		sinkSlice = res
	}
}

func BenchmarkCollectUnsized(b *testing.B) {
	data := generateData(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Map hides SizedIterator, forcing dynamic slice growth
		it := ops.Map(core.Slice(data), func(x int) int { return x })
		res := ops.Collect(it)
		sinkSlice = res
	}
}

func BenchmarkPipeline(b *testing.B) {
	data := generateData(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		filtered := ops.Filter(core.Slice(data), func(x int) bool { return x%2 == 0 })
		mapped := ops.Map(filtered, func(x int) int { return x * 2 })
		res := ops.Collect(mapped)
		sinkSlice = res
	}
}
