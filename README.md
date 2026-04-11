# _itFrame_

***itFrame*** is a **pull-based iterator framework** in Go designed for:

- lazy evaluation
- composable iteration
- minimal allocation overhead

---

### Features

- Generic Iterator interface
- SliceIterator
- RangeIterator (end-exclusive)
- MapIterator
- FilterIterator
- Terminal Operations (Reduce, Count, Any, All, Collect)
- Stream[T] abstraction over Iterators(chainable API on top of iterator system)
- Advanced Operations (Peek, Merge, Zip)
- Introduced ordering-aware operations (MergeDistinct, ZipWith, ZipLongest, Seek)
- FlatMap
- Relational Operations (Join, LeftJoin, Crossjoin, GroupBy) and Aggregation.
- Error Aware Iteration
- Added Benchmarks and edge-case tests
- Merged advanced into ops
- Added new features with some refinements

---

### Requirements

```go
go 1.23+
```

### Installation

```go
go get github.com/ver1619/itFrame
```

---

### To run tests

```go
go test ./...
```

---

### To run sample code

```go
go run ./examples/<file_path>
```
