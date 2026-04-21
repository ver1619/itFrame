# *stream*

The `stream` package gives you a **fluent, chainable API** for building data pipelines. Instead of nesting function calls like `ops.Filter(ops.Map(it, fn), pred)`, you can write `stream.Slice(data).Map(fn).Filter(pred).Collect()` — clean and readable.

Under the hood, `stream` wraps the same `ops` functions. Everything is still lazy — nothing runs until you call a terminal operation like `Collect()`, `Reduce()`, or `Count()`.

---

## Import

```go
import "github.com/ver1619/itFrame/stream"
```

---

## What's Inside

### Creating Streams
| Name | What It Does |
|------|-------------|
| `From` | Wraps any `core.Iterator[T]` into a Stream |
| `Slice` | Creates a Stream from a slice |
| `Range` | Creates a Stream of integers over `[start, end)` |

### Transform Operations (methods)
| Name | What It Does |
|------|-------------|
| `.Map(fn)` | Transforms each element (same type) |
| `.FlatMap(fn)` | Maps to slices and flattens (same type) |
| `.FlatMapIter(fn)` | Maps to iterators and flattens (same type) |
| `.Scan(init, fn)` | Emits running accumulations |

### Transform Operations (free functions — for type changes)
| Name | What It Does |
|------|-------------|
| `MapTo(s, fn)` | Transforms elements from type A to type B |
| `FlatMapTo(s, fn)` | Flat-maps with slices from A to B |
| `FlatMapIterTo(s, fn)` | Flat-maps with iterators from A to B |

### Filter & Control (methods)
| Name | What It Does |
|------|-------------|
| `.Filter(pred)` | Keeps elements matching a condition |
| `.Take(n)` | Yields at most N elements |
| `.Skip(n)` | Discards the first N elements |
| `.TakeWhile(pred)` | Yields while condition is true |
| `.DropWhile(pred)` | Skips while condition is true |
| `.Chain(other)` | Appends another stream |

### Terminal Operations (consume the stream)
| Name | What It Does |
|------|-------------|
| `.Collect()` | Returns all elements as a slice |
| `.Reduce(init, fn)` | Folds into a single value |
| `.Count()` | Returns the number of elements |
| `.ForEach(fn)` | Runs a function on each element |
| `.Any(pred)` | True if any element matches |
| `.All(pred)` | True if all elements match |

### Utility
| Name | What It Does |
|------|-------------|
| `.Iterator()` | Returns the underlying `core.Iterator[T]` |

---

## Syntax 

```go
// Create
s := stream.Slice([]int{1, 2, 3})
s := stream.Range(0, 10, 1)
s := stream.From(someIterator)

--------------------------------

// Chain operations (same type)
s.Map(fn).Filter(pred).Take(5).Collect()

--------------------------------

// Type-changing transforms (free functions)
stream.MapTo(s, fn)
stream.FlatMapTo(s, fn)
stream.FlatMapIterTo(s, fn)

--------------------------------

// Terminal
s.Collect()              // []T
s.Reduce(init, fn)       // T
s.Count()                // int
s.ForEach(fn)            // (nothing)
s.Any(pred)              // bool
s.All(pred)              // bool
```

> **Why free functions for type changes?** Go methods can't introduce new type parameters. So `Map` (T→T) is a method, but `MapTo` (A→B) must be a standalone function.

---

## Use Cases

- **Readable pipelines** — chain operations instead of nesting calls
- **Quick data processing** — filter, transform, and collect in one line
- **Prototyping** — try ideas faster with the fluent API
- **Pagination** — `.Skip(offset).Take(pageSize)`
- **Type conversions** — use `MapTo` to convert between types cleanly

---

## Examples

### Example 1 — Simple Filter-Map-Collect Pipeline

Filter even numbers and double them.

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/stream"
)

func main() {
    result := stream.Slice([]int{1, 2, 3, 4, 5, 6, 7, 8}).
        Filter(func(n int) bool { return n%2 == 0 }).
        Map(func(n int) int { return n * 2 }).
        Collect()

    fmt.Println(result)
}
```

**Output:** `[4 8 12 16]`

---

### Example 2 — Multi-Step Pipeline with Type Changes

Process a list of raw scores: drop any below 50, convert to letter grades, and collect.

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/stream"
)

func letterGrade(score int) string {
    switch {
    case score >= 90:
        return "A"
    case score >= 80:
        return "B"
    case score >= 70:
        return "C"
    case score >= 60:
        return "D"
    default:
        return "F"
    }
}

func main() {
    scores := stream.Slice([]int{95, 42, 78, 88, 55, 67, 91, 30})

    // Keep passing scores, then convert to letter grades
    passing := scores.Filter(func(n int) bool { return n >= 50 })
    grades := stream.MapTo(passing, letterGrade)

    fmt.Println("Grades:", grades.Collect())
    // Count how many passed (need a fresh stream since iterators are single-pass)
    count := stream.Slice([]int{95, 42, 78, 88, 55, 67, 91, 30}).
        Filter(func(n int) bool { return n >= 50 }).
        Count()
    fmt.Println("Passed:", count)

    // Running sum of first 4 elements
    sums := stream.Range(1, 6, 1).
        Scan(0, func(acc, val int) int { return acc + val }).
        Collect()
    fmt.Println("Running sums:", sums)

    // Check conditions
    allPositive := stream.Slice([]int{1, 2, 3}).
        All(func(n int) bool { return n > 0 })
    fmt.Println("All positive:", allPositive)

    hasNegative := stream.Slice([]int{1, -2, 3}).
        Any(func(n int) bool { return n < 0 })
    fmt.Println("Has negative:", hasNegative)
}
```

**Output:**
```
Grades: [A C B F D A]
Passed: 6
Running sums: [1 3 6 10 15]
All positive: true
Has negative: true
```
