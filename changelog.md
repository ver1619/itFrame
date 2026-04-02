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
