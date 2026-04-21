package benchmarks

import "testing"

var sink int
var sinkSlice []int

func generateData(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	return data
}

// ---------------- MAP ----------------

func BenchmarkMap(b *testing.B) {
	data := generateData(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum := 0
		for _, v := range data {
			sum += v * 2
		}
		sink = sum
	}
}

// ---------------- FILTER ----------------

func BenchmarkFilter(b *testing.B) {
	data := generateData(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum := 0
		for _, v := range data {
			if v%2 == 0 {
				sum += v
			}
		}
		sink = sum
	}
}

// ---------------- FLATMAP ----------------

func BenchmarkFlatMap(b *testing.B) {
	data := generateData(1000)
	inner := []int{1, 2, 3}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum := 0
		for range data {
			for _, v := range inner {
				sum += v
			}
		}
		sink = sum
	}
}

// ---------------- MERGE ----------------

func BenchmarkMerge(b *testing.B) {
	data1 := generateData(5000)
	data2 := generateData(5000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		i1, i2 := 0, 0
		sum := 0

		for i1 < len(data1) && i2 < len(data2) {
			if data1[i1] < data2[i2] {
				sum += data1[i1]
				i1++
			} else {
				sum += data2[i2]
				i2++
			}
		}
		for i1 < len(data1) {
			sum += data1[i1]
			i1++
		}
		for i2 < len(data2) {
			sum += data2[i2]
			i2++
		}

		sink = sum
	}
}

// ---------------- JOIN ----------------

func BenchmarkJoin(b *testing.B) {
	left := generateData(1000)
	right := generateData(500)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		hash := make(map[int][]int)

		for _, r := range right {
			k := r % 100
			hash[k] = append(hash[k], r)
		}

		count := 0
		for _, l := range left {
			k := l % 100
			if matches, ok := hash[k]; ok {
				count += len(matches)
			}
		}

		sink = count
	}
}

// ---------------- COLLECT (SIZED) ----------------

func BenchmarkCollectSized(b *testing.B) {
	data := generateData(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res := make([]int, 0, len(data))
		res = append(res, data...)
		sinkSlice = res
	}
}

// ---------------- COLLECT (UNSIZED) ----------------

func BenchmarkCollectUnsized(b *testing.B) {
	data := generateData(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res := []int{}
		res = append(res, data...)
		sinkSlice = res
	}
}

// ---------------- PIPELINE ----------------

func BenchmarkPipeline(b *testing.B) {
	data := generateData(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res := make([]int, 0, len(data)/2)
		for _, v := range data {
			if v%2 == 0 {
				res = append(res, v*2)
			}
		}
		sinkSlice = res
	}
}
