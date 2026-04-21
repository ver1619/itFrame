# *core*

The `core` package is the foundation of itFrame. It gives you the basic building blocks — the `Iterator` interface that everything else is built on, and simple ways to create iterators from slices and number ranges.

Think of it like this: before you can filter, map, or transform data, you need a way to *walk through* it one item at a time. That's what `core` does.

---

## Import

```go
import "github.com/ver1619/itFrame/core"
```

---

## What's Inside

| Name | What It Does |
|------|-------------|
| `Iterator[T]` | The interface every iterator uses — just one method: `Next()` |
| `Slice[T]` | Turns a slice into an iterator |
| `Range` | Creates an iterator over a range of integers |
| `SizedIterator` | Optional interface for iterators that know their length |

---

## Syntax

### Iterator Interface

```go
type Iterator[T any] interface {
    Next() (T, bool)
}
```

Call `Next()` to get the next item. It returns the value and `true`, or a zero value and `false` when there's nothing left.

### Slice

```go
core.Slice[T](data []T) *SliceIterator[T]
```

### Range

```go
core.Range(start, end, step int) *RangeIterator
```

Creates numbers from `start` up to (but not including) `end`, stepping by `step`. Panics if `step` is zero.

---

## Use Cases

- **Starting point for every pipeline** — you always begin with a `core` source
- **Iterating over slices** without copying them
- **Generating number sequences** for loops, indices, or test data
- **Building custom iterators** by implementing the `Iterator[T]` interface

---

## Examples

### Example 1 — Iterate Over a Slice

Walk through a list of names and print each one.

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
)

func main() {
    names := []string{"Alice", "Bob", "Charlie"}
    it := core.Slice(names)

    for name, ok := it.Next(); ok; name, ok = it.Next() {
        fmt.Println(name)
    }
}
```

**Output:**
```
Alice
Bob
Charlie
```

---

### Example 2 — Generate a Number Range and Check Length

Create a countdown from 10 to 1 using a negative step, and use `Len()` to check how many items remain as you go.

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
)

func main() {
    countdown := core.Range(10, 0, -1)

    fmt.Println("Total items:", countdown.Len()) // 10

    // Pull the first 3 values
    for i := 0; i < 3; i++ {
        val, _ := countdown.Next()
        fmt.Println("Value:", val)
    }

    fmt.Println("Remaining:", countdown.Len()) // 7

    // Create a slice iterator and check its size too
    scores := core.Slice([]int{90, 85, 72, 95})
    fmt.Println("Scores count:", scores.Len()) // 4

    scores.Next() // consume one
    fmt.Println("After one Next():", scores.Len()) // 3
}
```

**Output:**
```
Total items: 10
Value: 10
Value: 9
Value: 8
Remaining: 7
Scores count: 4
After one Next(): 3
```
