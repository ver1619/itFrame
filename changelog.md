# Initial Release

## v0.1.0 - 2026-04-01

### Added

- `Iterator[T]` interface (pull-based contract)
- `SliceIterator`
- `RangeIterator` (end-exclusive semantics)

### Notes

- Lazy, single-pass iteration model

---

# Second Release

## v0.2.0 - 2026-04-02

### Added

- `MapIterator`
- `FilterIterator`

### Modified

- Changed function naming<br>
  `NewSliceIterator` -> `Slice`<br>
  `NewRangeIterator` -> `Range`

### Notes

- Lazy transformation and filtering
- Composable iterator pipeline introduced

---

# Third Release

## v0.3.0 - 2026-04-03

### Added

- `Reduce`
- `Count`
- `Collect`
- `Any` / `All`

### Modified
- Refined and formatted comments in `./core` `./ops`

### Notes
- Introduced **Terminal operations**
- Iterators are consumed after terminal execution
- `Any` / `All` support short-circuit evaluation

---

# Fourth Release

## v0.4.0 - 2026-04-04

###  Added

- `Stream[T]` abstraction over Iterator<br>
- Fluent API for:
    - Map
    - Filter
- Terminal methods on Stream (`Reduce`, `Collect`, `Count`, `Any`, `All`)    


### Notes
- Introduced `Stream` as the primary user-facing API
- `Stream` wraps underlying iterators and preserves lazy evaluation
- Supports composable, chainable pipelines
- Terminal operations consume the stream
- Existing iterator-based operations remain as lower-level primitives

---

# Fifth Release

## v0.5.0 - 2026-04-05

### Added

- `PeekIterator`
- `MergeIterator` (stable merge)
- `ZipIterator`

### Notes

- Introduced multi-source iteration
- Added lookahead capability
- Supports sorted merge pipelines

---

# Sixth Release

## v0.6.0 - 2026-04-06

### Added
- `Comparator abstraction`
- `MergeDistinctIterator`
- `SeekIterator`
- `ZipWithIterator`
- `ZipLongestIterator`

### Notes
- Introduced ordering-aware operations
- Enabled deduplication and alignment semantics
- Prepared foundation for relational operations (v0.7)

---

# Seventh Release

## v0.7.0 - 2026-04-07

### Added
- `FlatMap`
- `CrossJoin`
- `Join` (inner join)
- `LeftJoin`
- `GroupBy`
- `Aggregate`

### Notes
- Introduced relational operations
- Enabled grouping and aggregation
- Transitioned toward query-engine capabilities

---

# Eighth Release

## v0.8.0 - 2026-04-08

### Added
- `Result[T]` for error-aware values
- Error-aware `Map`, `Filter`, `FlatMap`
- Error-aware Stream API
- `Collect` with error handling

### Notes
- Introduced error propagation in pipelines
- Enabled safe data processing workflows

---

# Ninth Release

## v0.9.0 - 2026-04-13

### Added

- Core control operations:
  - `Take(n)`, `Skip(n)`
  - `TakeWhile`, `DropWhile`
  - `Chain(it1, it2)`

- Terminal operation:
  - `ForEach`

- Advanced operations:
  - `Scan` (running accumulation)
  - `GroupBySorted` (memory-efficient grouping for sorted inputs)
  - `Distinct` (comparator-based de-duplication)

- `FlatMapIter` for iterator-based expansion
- Benchmark suite for performance comparison (native vs itFrame)

### Modified

- `FlatMap` redesigned to use slice-based expansion  (`func(T) []U`)
- Previous iterator-based behavior moved to `FlatMapIter`
- Stream API extended with new control and terminal operations
- Unified API surface by merging `advanced` operations into `ops`
- Examples and tests updated to reflect new APIs

### Fixed

- Resolved all `go vet` issues
- Removed unused/dead internal code
- Improved documentation with full godoc coverage

### Notes
- `FlatMap` is now optimized for performance and is the default choice
- `FlatMapIter` remains available for iterator-based use-cases
- API surface simplified with all operations under `ops`
- Benchmarks and edge-case tests added to validate performance and correctness