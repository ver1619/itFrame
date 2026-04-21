# _itFrame_

A pull-based, generic iterator framework for Go — build lazy, composable, and allocation-efficient data pipelines.

[![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/ver1619/itFrame)](https://goreportcard.com/report/github.com/ver1619/itFrame)
[![CI](https://github.com/ver1619/itFrame/actions/workflows/ci.yml/badge.svg)](https://github.com/ver1619/itFrame/actions/workflows/ci.yml)

---

## Overview

**itFrame** gives you a simple way to process data step-by-step in Go — without loading everything into memory at once. You describe *what* to do (filter, map, group, join) and itFrame figures out *when* to do it, pulling values through the pipeline only as needed.

Think of it like a conveyor belt: each item moves through your operations one at a time, and nothing runs until you ask for the final result.

---

## Key Features

- **Lazy Evaluation** — nothing executes until a terminal operation (`Collect`, `Reduce`, `Count`) is called
- **30+ Operations** — Map, Filter, FlatMap, Reduce, Scan, Take, Skip, Zip, Merge, Join, GroupBy, Aggregate, and more
- **Fluent Stream API** — chain operations with `stream.Slice(data).Filter(...).Map(...).Collect()`
- **Error-Aware Pipelines** — wrap values in `Result[T]` and errors propagate automatically without crashing
- **Zero-Allocation Core** — Map, Filter, FlatMap, and Merge run with **0 B/op** in native benchmarks
- **Relational Operations** — SQL-style `Join`, `LeftJoin`, `CrossJoin`, and `GroupBy` with `Aggregate`
- **Comparator System** — pluggable ordering for Merge, Distinct, and multi-field sorting
- **Pre-Allocation** — `SizedIterator` lets `Collect` pre-allocate slices for known-length sources

---

## Installation

### Requirements

```
Go 1.23 or higher
```

### Get itFrame

```bash
go get github.com/ver1619/itFrame
```

### Quick Start

Create a file `main.go`:

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/stream"
)

func main() {
    result := stream.Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
        Filter(func(n int) bool { return n%2 == 0 }).
        Map(func(n int) int { return n * n }).
        Collect()

    fmt.Println(result) // [4 16 36 64 100]
}
```

Run:

```bash
go run main.go
```

---

## Tests

itFrame has unit tests, edge-case tests, and race-condition tests across all packages. Tests cover normal operations, empty iterators, single-element inputs, large datasets, and concurrent safety.

### Run all tests

```bash
go test ./...
```

### Run with verbose output and race detection

```bash
go test -v -race ./...
```

---

## Core Concepts

### Iterator

The `Iterator[T]` interface is the foundation of everything in itFrame. It has a single method:

```go
type Iterator[T any] interface {
    Next() (T, bool)
}
```

Call `Next()` to get the next value. It returns `(value, true)` when there's data, or `(zero, false)` when the iterator is done. Every operation in itFrame either takes an iterator as input or returns one as output.

### Lazy Evaluation

Operations like `Map`, `Filter`, and `FlatMap` don't process any data when you create them — they just build a pipeline description. The actual work only happens when you call a **terminal operation** like `Collect()`, `Reduce()`, or `Count()`. This means:

- You only process what you need (use `Take(5)` to stop after 5 items)
- Memory stays low because items flow through one at a time
- You can build complex pipelines without intermediate slices

### Sized Iterator

Some iterators (like `SliceIterator` and `RangeIterator`) know exactly how many elements they'll produce. They implement the `SizedIterator` interface:

```go
type SizedIterator interface {
    Len() int
}
```

When `Collect()` detects a `SizedIterator`, it pre-allocates the result slice to the exact size — avoiding repeated slice growth and reducing allocations from **19 allocs** down to **1 alloc**.

---

## Documentation

### Package Guides

Each guide includes explanation, syntax, use cases, and runnable examples.

| Package | Description | Guide |
|---------|-------------|-------|
| **core** | Iterator interface, Slice and Range sources | [core.md](./package/core.md) |
| **ops** | 30+ standalone operations (Map, Filter, Join, GroupBy, etc.) | [ops.md](./package/ops.md) |
| **stream** | Fluent chainable API over iterators | [stream.md](./package/stream.md) |
| **errors** | Error-aware pipelines with `Result[T]` | [errors.md](./package/errors.md) |
| **compare** | Comparators for ordering, sorting, and equality | [compare.md](./package/compare.md) |

### Real-World Programs

Full programs with explanation, code, verified output, and architecture diagrams.

| Program | What It Demonstrates | Guide |
|---------|---------------------|-------|
| **Server Log Analyzer** | Parse → Filter → GroupBy → Aggregate → Merge | [log_analyzer.md](./examples/code_files/log_analyzer/log_analyzer.md) |
| **E-Commerce Analytics** | Join → Map → GroupBy → Aggregate → Reduce + pagination | [e_commerce_analytics.md](./examples/code_files/e_commerce_analytics/e_commerce_analytics.md) |
| **Sensor Data Pipeline** | Error handling → MergeDistinct → Reduce → Scan | [sensor_pipeline.md](./examples/code_files/sensor_pipeline/sensor_pipeline.md) |

---

## Design Decisions

- **Pull-based, not push-based** — the consumer controls the pace by calling `Next()`, making it easy to stop early, compose operations, and avoid buffering data you don't need
- **Free functions for type-changing operations** — Go methods cannot introduce new type parameters, so operations like `MapTo[A, B]` are standalone functions while same-type operations like `Map[T]` are methods on `Stream`
- **Slice-based FlatMap alongside iterator-based** — `FlatMap(fn func(A) []B)` avoids the per-element interface allocation that `FlatMapIter(fn func(A) Iterator[B])` incurs, giving a 4-5× performance boost for the common case
- **Comparator interface instead of `<` operator** — a `Less(a, b) bool` interface supports custom types, multi-field sorting via `Multi`, and reverse ordering via `Reverse` without code duplication
- **Errors as values, not panics** — the `errors` package uses `Result[T]` to wrap values or errors, letting pipelines propagate failures automatically without `if err != nil` at every step

---

## Benchmarks

Benchmarks compare **itFrame operations** against **native Go loops** on 10,000 elements.

### Results

| Operation | itFrame | Native Go | Overhead | itFrame Allocs |
|-----------|---------|-----------|----------|----------------|
| Map | 124,961 ns/op | 5,909 ns/op | ~21× | 2 allocs |
| Filter | 106,152 ns/op | 10,603 ns/op | ~10× | 2 allocs |
| FlatMap (slice) | 28,371 ns/op | 4,813 ns/op | ~6× | 3 allocs |
| FlatMap (iter) | 132,536 ns/op | 4,813 ns/op | ~28× | 1,003 allocs |
| Merge | 248,032 ns/op | 17,067 ns/op | ~15× | 3 allocs |
| Join | 115,443 ns/op | 68,410 ns/op | ~1.7× | 414 allocs |
| Collect (sized) | 95,290 ns/op | 53,678 ns/op | ~1.8× | 2 allocs |
| Collect (unsized) | 318,118 ns/op | 189,236 ns/op | ~1.7× | 21 allocs |
| Pipeline | 211,111 ns/op | 34,412 ns/op | ~6× | 19 allocs |

### Insights

- **FlatMap (slice) vs FlatMap (iter)** — slice-based FlatMap is **4.7× faster** and uses **334× fewer allocations** because it avoids creating a new iterator object per element
- **Sized vs Unsized Collect** — pre-allocation cuts collection time by **3.3×** and allocations from 21 to 2
- **Join is near-native** — only 1.7× overhead since both approaches use hash-based lookup
- **Core transforms (Map, Filter)** have higher relative overhead due to interface dispatch on `Next()` calls, but the absolute cost is still microseconds for 10K elements

### Run benchmarks

```bash
# itFrame package benchmarks
go test -bench=. -benchmem ./benchmarks/ -run='^$' -args -mode=package

# Native Go baseline benchmarks
go test -bench=. -benchmem ./benchmarks/ -run='^$' -args -mode=native
```

---

## Limitations

- **Single-pass only** — iterators can only be consumed once. If you need the data again, call `Collect()` first and create a new iterator from the slice
- **No parallel execution** — all operations run sequentially on a single goroutine. itFrame is not designed for concurrent fan-out pipelines
- **Interface overhead on hot paths** — the `Iterator[T]` interface adds virtual dispatch cost on every `Next()` call, which matters in tight loops over millions of elements
- **GroupBy buffers everything** — `GroupBy` must consume the entire input before yielding groups. For sorted data, use `GroupBySorted` which streams lazily with O(group-size) memory
- **Type-changing operations require free functions** — due to Go's generics limitations, `MapTo`, `FlatMapTo`, and `FlatMapIterTo` must be called as standalone functions instead of methods, which breaks the fluent chain

---

## License

[MIT](./LICENSE) © Swayam Vernekar
