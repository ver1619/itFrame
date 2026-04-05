# Initial Release

## v0.1.0 - 2026-04-01

### Added

- Iterator[T] interface (pull-based contract)
- SliceIterator
- RangeIterator (end-exclusive semantics)

### Notes

- Lazy, single-pass iteration model

---

# Second Release

## v0.2.0 - 2026-04-02

### Added

- MapIterator
- FilterIterator

### Modify

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

### Modify
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
- Terminal methods on Stream (Reduce, Collect, Count, Any, All)    


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

- PeekIterator
- MergeIterator (stable merge)
- ZipIterator

### Notes

- Introduced multi-source iteration
- Added lookahead capability
- Supports sorted merge pipelines

---
