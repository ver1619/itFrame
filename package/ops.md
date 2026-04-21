# *ops*

The `ops` package is the toolbox of itFrame. It has all the operations you need to transform, filter, combine, and consume iterators. Every function here takes an iterator in and gives you either a new iterator or a final result.

These are **standalone functions** — you pass iterators to them directly. If you prefer method chaining, check out the `stream` package which wraps these into a fluent API.

---

## Import

```go
import "github.com/ver1619/itFrame/ops"
```

---

## What's Inside

### Transform Operations
| Name | What It Does |
|------|-------------|
| `Map` | Transforms each element using a function |
| `FlatMap` | Maps each element to a slice and flattens the results |
| `FlatMapIter` | Maps each element to an iterator and flattens |
| `Scan` | Emits a running accumulation after each element |

### Filter & Control Operations
| Name | What It Does |
|------|-------------|
| `Filter` | Keeps only elements that match a condition |
| `Take` | Yields at most the first N elements |
| `Skip` | Discards the first N elements |
| `TakeWhile` | Yields elements while a condition is true |
| `DropWhile` | Skips elements while a condition is true, then yields the rest |
| `Seek` | Skips until a condition is true, then yields from there |
| `Distinct` | Removes consecutive duplicates from sorted data |
| `Peek` | Lets you look at the next element without consuming it |

### Combine Operations
| Name | What It Does |
|------|-------------|
| `Chain` | Concatenates two iterators end-to-end |
| `Zip` | Pairs elements from two iterators (stops at shorter) |
| `ZipLongest` | Pairs elements from two iterators (continues to longer) |
| `ZipWith` | Combines two iterators using a custom function |
| `Merge` | Merges two sorted iterators into one sorted sequence |
| `MergeDistinct` | Merges two sorted iterators, removing duplicates |
| `CrossJoin` | Produces every possible pair from two iterators |

### Join Operations
| Name | What It Does |
|------|-------------|
| `Join` | Inner join — only matching pairs are returned |
| `LeftJoin` | Left join — all left elements, matched or not |

### Group Operations
| Name | What It Does |
|------|-------------|
| `GroupBy` | Groups elements by a key (buffers everything) |
| `GroupBySorted` | Groups pre-sorted elements lazily (memory efficient) |
| `Aggregate` | Summarizes each group into a single result |

### Terminal Operations (consume the iterator)
| Name | What It Does |
|------|-------------|
| `Collect` | Gathers all elements into a slice |
| `Reduce` | Folds all elements into one value |
| `Count` | Returns the number of elements |
| `ForEach` | Runs a function on each element |
| `Any` | True if at least one element matches |
| `All` | True if every element matches |

### Types
| Name | What It Does |
|------|-------------|
| `Pair[A, B]` | A generic tuple with `First` and `Second` fields |
| `Group[K, V]` | A key-value group with `Key` and `Items` fields |

---

## Syntax

```go
// Transform
ops.Map[A, B](it, fn)
ops.FlatMap[A, B](it, fn)       // fn returns []B
ops.FlatMapIter[A, B](it, fn)   // fn returns Iterator[B]
ops.Scan[T, R](it, init, fn)

---------------------------------

// Filter & Control
ops.Filter[T](it, pred)
ops.Take[T](it, n)
ops.Skip[T](it, n)
ops.TakeWhile[T](it, pred)
ops.DropWhile[T](it, pred)
ops.Seek[T](it, pred)
ops.Distinct[T](it, cmp)
ops.Peek[T](it)

---------------------------------

// Combine
ops.Chain[T](it1, it2)
ops.Zip[A, B](it1, it2)
ops.ZipLongest[A, B](it1, it2)
ops.ZipWith[A, B, C](it1, it2, fn)
ops.Merge[T](it1, it2, cmp)
ops.MergeDistinct[T](it1, it2, cmp)
ops.CrossJoin[A, B](it1, it2)

---------------------------------

// Joins
ops.Join[A, B, K](left, right, leftKey, rightKey)
ops.LeftJoin[A, B, K](left, right, leftKey, rightKey)

---------------------------------

// Group
ops.GroupBy[T, K](it, keyFn)
ops.GroupBySorted[T, K](it, keyFn)
ops.Aggregate[K, V, R](it, fn)

---------------------------------

// Terminal
ops.Collect[T](it)
ops.Reduce[T, R](it, init, fn)
ops.Count[T](it)
ops.ForEach[T](it, fn)
ops.Any[T](it, pred)
ops.All[T](it, pred)
```

---

## Use Cases

- **Data transformation pipelines** — map, filter, and collect in sequence
- **Pagination** — use Skip + Take to grab a page of results
- **Joining datasets** — combine related data like SQL joins
- **Aggregation and reporting** — group data and summarize each group
- **Stream merging** — combine sorted streams from multiple sources
- **Deduplication** — remove duplicates from sorted data
- **Running totals** — compute cumulative sums with Scan
- **Validation** — check if Any or All elements meet a condition

---

## Examples

Since `ops` has many operations, here are **2 examples per category**.

---

### Map

#### Example 1 — Double Every Number

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    it := core.Slice([]int{1, 2, 3, 4})
    doubled := ops.Map(it, func(n int) int { return n * 2 })
    fmt.Println(ops.Collect(doubled))
}
```

**Output:** `[2 4 6 8]`

#### Example 2 — Convert Structs to Strings

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

type User struct {
    Name string
    Age  int
}

func main() {
    users := core.Slice([]User{
        {"Alice", 30},
        {"Bob", 25},
    })

    labels := ops.Map(users, func(u User) string {
        return fmt.Sprintf("%s (age %d)", u.Name, u.Age)
    })

    ops.ForEach(labels, func(s string) { fmt.Println(s) })
}
```

**Output:**
```
Alice (age 30)
Bob (age 25)
```

---

### Filter

#### Example 1 — Keep Even Numbers

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    it := core.Slice([]int{1, 2, 3, 4, 5, 6})
    evens := ops.Filter(it, func(n int) bool { return n%2 == 0 })
    fmt.Println(ops.Collect(evens))
}
```

**Output:** `[2 4 6]`

#### Example 2 — Filter Structs by Field

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

type Product struct {
    Name  string
    Price float64
}

func main() {
    products := core.Slice([]Product{
        {"Laptop", 999.99},
        {"Mouse", 29.99},
        {"Monitor", 449.99},
        {"Cable", 9.99},
    })

    expensive := ops.Filter(products, func(p Product) bool {
        return p.Price > 100
    })

    ops.ForEach(expensive, func(p Product) {
        fmt.Printf("%s: $%.2f\n", p.Name, p.Price)
    })
}
```

**Output:**
```
Laptop: $999.99
Monitor: $449.99
```

---

### FlatMap / FlatMapIter

#### Example 1 — Expand Each Number into a Range

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    it := core.Slice([]int{1, 2, 3})

    // Each number n produces [1, 2, ..., n]
    expanded := ops.FlatMap(it, func(n int) []int {
        result := make([]int, n)
        for i := range result {
            result[i] = i + 1
        }
        return result
    })

    fmt.Println(ops.Collect(expanded))
}
```

**Output:** `[1 1 2 1 2 3]`

#### Example 2 — FlatMapIter with Iterator-Producing Function

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    words := core.Slice([]string{"hello", "go"})

    // For each word, create an iterator over its characters
    chars := ops.FlatMapIter(words, func(w string) core.Iterator[byte] {
        return core.Slice([]byte(w))
    })

    ops.ForEach(chars, func(b byte) {
        fmt.Printf("%c ", b)
    })
}
```

**Output:** `h e l l o g o`

---

### Take / Skip / TakeWhile / DropWhile

#### Example 1 — Paginate with Skip and Take

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    data := core.Slice([]string{"a", "b", "c", "d", "e", "f", "g"})

    // Get page 2 (items 3-4), page size = 2
    page := ops.Take(ops.Skip(data, 2), 2)
    fmt.Println(ops.Collect(page))
}
```

**Output:** `[c d]`

#### Example 2 — TakeWhile and DropWhile on Sorted Data

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    // Sorted scores
    scores := []int{45, 52, 67, 78, 85, 91, 95}

    // Take scores below 80
    below80 := ops.TakeWhile(core.Slice(scores), func(n int) bool { return n < 80 })
    fmt.Println("Below 80:", ops.Collect(below80))

    // Drop scores below 80, get the rest
    above80 := ops.DropWhile(core.Slice(scores), func(n int) bool { return n < 80 })
    fmt.Println("80 and above:", ops.Collect(above80))
}
```

**Output:**
```
Below 80: [45 52 67 78]
80 and above: [78 85 91 95]
```

---

### Seek / Peek

#### Example 1 — Seek to a Starting Point

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    logs := core.Slice([]string{"debug: init", "debug: load", "info: ready", "info: serving", "error: timeout"})

    // Skip all debug logs, start from the first non-debug entry
    fromInfo := ops.Seek(logs, func(s string) bool {
        return s[:4] != "debu"
    })

    fmt.Println(ops.Collect(fromInfo))
}
```

**Output:** `[info: ready info: serving error: timeout]`

#### Example 2 — Peek Without Consuming

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    it := ops.Peek(core.Slice([]int{10, 20, 30}))

    // Peek doesn't consume
    val, _ := it.Peek()
    fmt.Println("Peeked:", val)

    // Peek again — same value
    val, _ = it.Peek()
    fmt.Println("Peeked again:", val)

    // Now consume
    val, _ = it.Next()
    fmt.Println("Consumed:", val)

    // Next peek shows the new head
    val, _ = it.Peek()
    fmt.Println("New peek:", val)
}
```

**Output:**
```
Peeked: 10
Peeked again: 10
Consumed: 10
New peek: 20
```

---

### Chain / Zip / ZipWith / ZipLongest

#### Example 1 — Chain Two Lists Together

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    first := core.Slice([]int{1, 2, 3})
    second := core.Slice([]int{4, 5, 6})

    combined := ops.Chain(first, second)
    fmt.Println(ops.Collect(combined))
}
```

**Output:** `[1 2 3 4 5 6]`

#### Example 2 — ZipWith to Combine Two Lists

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    names := core.Slice([]string{"Alice", "Bob", "Charlie"})
    scores := core.Slice([]int{95, 87, 72})

    // Combine into formatted strings
    results := ops.ZipWith(names, scores, func(name string, score int) string {
        return fmt.Sprintf("%s: %d", name, score)
    })

    ops.ForEach(results, func(s string) { fmt.Println(s) })
}
```

**Output:**
```
Alice: 95
Bob: 87
Charlie: 72
```

---

### Merge / MergeDistinct / Distinct

#### Example 1 — Merge Two Sorted Lists

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/compare"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    cmp := compare.LessFunc[int](func(a, b int) bool { return a < b })

    a := core.Slice([]int{1, 3, 5, 7})
    b := core.Slice([]int{2, 4, 6, 8})

    merged := ops.Merge(a, b, cmp)
    fmt.Println(ops.Collect(merged))
}
```

**Output:** `[1 2 3 4 5 6 7 8]`

#### Example 2 — MergeDistinct to Deduplicate

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/compare"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    cmp := compare.LessFunc[int](func(a, b int) bool { return a < b })

    list1 := core.Slice([]int{1, 2, 3, 5})
    list2 := core.Slice([]int{2, 3, 4, 6})

    unique := ops.MergeDistinct(list1, list2, cmp)
    fmt.Println(ops.Collect(unique))
}
```

**Output:** `[1 2 3 4 5 6]`

---

### Join / LeftJoin / CrossJoin

#### Example 1 — Inner Join Two Datasets

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    type Order struct {
        ID     int
        UserID int
        Item   string
    }
    type User struct {
        ID   int
        Name string
    }

    users := core.Slice([]User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}})
    orders := core.Slice([]Order{{101, 1, "Laptop"}, {102, 2, "Mouse"}, {103, 1, "Monitor"}})

    joined := ops.Join(
        users, orders,
        func(u User) int { return u.ID },
        func(o Order) int { return o.UserID },
    )

    ops.ForEach(joined, func(p ops.Pair[User, Order]) {
        fmt.Printf("%s ordered %s\n", p.First.Name, p.Second.Item)
    })
}
```

**Output:**
```
Alice ordered Laptop
Alice ordered Monitor
Bob ordered Mouse
```

#### Example 2 — LeftJoin (Keeps Unmatched Left Items)

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    type Student struct {
        ID   int
        Name string
    }
    type Grade struct {
        StudentID int
        Score     int
    }

    students := core.Slice([]Student{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}})
    grades := core.Slice([]Grade{{1, 95}, {3, 78}})

    result := ops.LeftJoin(
        students, grades,
        func(s Student) int { return s.ID },
        func(g Grade) int { return g.StudentID },
    )

    ops.ForEach(result, func(p ops.Pair[Student, Grade]) {
        if p.Second.Score == 0 {
            fmt.Printf("%s: no grade submitted\n", p.First.Name)
        } else {
            fmt.Printf("%s: %d\n", p.First.Name, p.Second.Score)
        }
    })
}
```

**Output:**
```
Alice: 95
Bob: no grade submitted
Charlie: 78
```

---

### GroupBy / GroupBySorted / Aggregate

#### Example 1 — Group Words by First Letter

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    words := core.Slice([]string{"apple", "ant", "banana", "avocado", "berry"})

    groups := ops.GroupBy(words, func(w string) byte { return w[0] })

    ops.ForEach(groups, func(g ops.Group[byte, string]) {
        fmt.Printf("%c: %v\n", g.Key, g.Items)
    })
}
```

**Output (order may vary):**
```
a: [apple ant avocado]
b: [banana berry]
```

#### Example 2 — GroupBySorted + Aggregate for Summaries

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

type Sale struct {
    Dept   string
    Amount int
}

func main() {
    // Data must be pre-sorted by key for GroupBySorted
    sales := core.Slice([]Sale{
        {"Engineering", 500},
        {"Engineering", 300},
        {"Marketing", 200},
        {"Marketing", 400},
        {"Marketing", 100},
    })

    groups := ops.GroupBySorted(sales, func(s Sale) string { return s.Dept })

    summaries := ops.Aggregate(groups, func(g ops.Group[string, Sale]) string {
        total := 0
        for _, s := range g.Items {
            total += s.Amount
        }
        return fmt.Sprintf("%s: $%d (%d sales)", g.Key, total, len(g.Items))
    })

    ops.ForEach(summaries, func(s string) { fmt.Println(s) })
}
```

**Output:**
```
Engineering: $800 (2 sales)
Marketing: $700 (3 sales)
```

---

### Scan / Reduce

#### Example 1 — Sum with Reduce

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    it := core.Slice([]int{10, 20, 30, 40})
    total := ops.Reduce(it, 0, func(acc, val int) int { return acc + val })
    fmt.Println("Total:", total)
}
```

**Output:** `Total: 100`

#### Example 2 — Running Total with Scan

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    it := core.Slice([]int{10, 20, 30, 40})
    running := ops.Scan(it, 0, func(acc, val int) int { return acc + val })
    fmt.Println(ops.Collect(running))
}
```

**Output:** `[10 30 60 100]`

---

### Collect / Count / ForEach / Any / All

#### Example 1 — Collect and Count

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    it := core.Slice([]int{5, 10, 15})
    result := ops.Collect(it)
    fmt.Println("Items:", result)

    count := ops.Count(core.Slice([]int{5, 10, 15}))
    fmt.Println("Count:", count)
}
```

**Output:**
```
Items: [5 10 15]
Count: 3
```

#### Example 2 — Any and All Checks

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/core"
    "github.com/ver1619/itFrame/ops"
)

func main() {
    ages := []int{22, 17, 30, 15, 45}

    hasMinor := ops.Any(core.Slice(ages), func(a int) bool { return a < 18 })
    fmt.Println("Has minor:", hasMinor)

    allAdults := ops.All(core.Slice(ages), func(a int) bool { return a >= 18 })
    fmt.Println("All adults:", allAdults)
}
```

**Output:**
```
Has minor: true
All adults: false
```
