# *errors*

The `errors` package lets you build data pipelines that **handle errors gracefully**. Instead of checking `err != nil` after every step, you wrap values in a `Result[T]` and errors flow through the pipeline automatically — just like the happy-path values.

If an error appears, transform operations like `Map` and `Filter` skip the bad item and pass the error along untouched. When you `Collect` at the end, the first error stops everything and gets returned.

---

## Import

```go
import "github.com/ver1619/itFrame/errors"
```

---

## What's Inside

### Result Type
| Name | What It Does |
|------|-------------|
| `Result[T]` | Wraps a value **or** an error — has `Value` and `Err` fields |
| `Ok(v)` | Creates a successful result |
| `ErrResult(err)` | Creates an error result |
| `IsError()` | Returns true if the result holds an error |

### Standalone Operations (function-style)
| Name | What It Does |
|------|-------------|
| `Map` | Transforms valid values, passes errors through |
| `Filter` | Keeps valid values matching a condition, passes errors through |
| `Collect` | Gathers valid values into a slice, stops on the first error |

### Stream API (method chaining)
| Name | What It Does |
|------|-------------|
| `FromIterator` | Creates a Stream from an iterator of Result values |
| `FromSlice` | Creates a Stream from a slice of Result values |
| `.Map(fn)` | Transforms valid values |
| `.Filter(pred)` | Filters valid values |
| `.FlatMap(fn)` | Flat-maps valid values to iterators |
| `.Collect()` | Returns `([]T, error)` |
| `.Iterator()` | Returns the underlying iterator |

---

## Syntax

### Creating Results

```go
good := errors.Ok(42)
bad  := errors.ErrResult[int](fmt.Errorf("something went wrong"))

good.IsError() // false
bad.IsError()  // true
```

### Function-Style Pipeline

```go
mapped  := errors.Map(it, func(v int) int { return v * 2 })
filtered := errors.Filter(it, func(v int) bool { return v > 10 })
values, err := errors.Collect(it)
```

### Stream-Style Pipeline

```go
values, err := errors.FromSlice(data).
    Map(fn).
    Filter(pred).
    Collect()
```

---

## Use Cases

- **Parsing pipelines** — parse strings to numbers, propagate parse errors automatically
- **I/O processing** — read lines from a file, skip bad ones, collect good ones
- **Validation** — check each item, mark invalid ones as errors
- **API responses** — process results from external calls where some may fail
- **Data cleaning** — transform data where some entries are malformed

---

## Examples

### Example 1 — Parse Numbers with Error Handling

Parse a list of strings to integers. Invalid strings become errors that propagate through the pipeline.

```go
package main

import (
    "fmt"
    "strconv"

    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/errors"
)

func main() {
    raw := []string{"10", "20", "abc", "40"}

    // Convert strings to Result[int]
    results := make([]errors.Result[int], len(raw))
    for i, s := range raw {
        n, err := strconv.Atoi(s)
        if err != nil {
            results[i] = errors.ErrResult[int](fmt.Errorf("bad value: %s", s))
        } else {
            results[i] = errors.Ok(n)
        }
    }

    it := core.Slice(results)

    // Double the valid numbers — errors pass through automatically
    doubled := errors.Map(it, func(n int) int { return n * 2 })

    values, err := errors.Collect(doubled)
    if err != nil {
        fmt.Println("Stopped on error:", err)
    } else {
        fmt.Println("Values:", values)
    }
}
```

**Output:** `Stopped on error: bad value: abc`

---

### Example 2 — Stream Pipeline with Mixed Results

Use the fluent Stream API to filter and transform data that includes errors, collecting only when everything is valid.

```go
package main

import (
    "fmt"

    "github.com/ver1619/itFrame/errors"
)

type Measurement struct {
    Sensor string
    Value  float64
}

func main() {
    // Some readings are valid, one has an error
    readings := []errors.Result[Measurement]{
        errors.Ok(Measurement{"temp", 22.5}),
        errors.Ok(Measurement{"temp", 25.1}),
        errors.Ok(Measurement{"humidity", 60.0}),
        errors.ErrResult[Measurement](fmt.Errorf("sensor disconnected")),
        errors.Ok(Measurement{"temp", 23.8}),
    }

    // Build a pipeline: keep only temp readings, transform to Celsius label
    values, err := errors.FromSlice(readings).
        Filter(func(m Measurement) bool { return m.Sensor == "temp" }).
        Map(func(m Measurement) Measurement {
            m.Value = m.Value + 273.15 // Convert to Kelvin
            return m
        }).
        Collect()

    if err != nil {
        fmt.Println("Pipeline error:", err)
    } else {
        for _, v := range values {
            fmt.Printf("%s: %.2f K\n", v.Sensor, v.Value)
        }
    }

    // Now try with all-valid data
    goodReadings := []errors.Result[int]{
        errors.Ok(10),
        errors.Ok(20),
        errors.Ok(30),
    }

    vals, err := errors.FromSlice(goodReadings).
        Filter(func(n int) bool { return n > 15 }).
        Map(func(n int) int { return n * 100 }).
        Collect()

    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Processed:", vals)
    }
}
```

**Output:**
```
Pipeline error: sensor disconnected
Processed: [2000 3000]
```
