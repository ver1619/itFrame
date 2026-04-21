# *compare*

The `compare` package gives you a clean way to define how things are ordered. Instead of scattering `<` comparisons everywhere, you define a `Comparator` once and reuse it across sorting, merging, deduplication, and more.

It also comes with handy tools — reverse the order, combine multiple sort criteria, or check equality — all built on a single `Less` method.

---

## Import

```go
import "github.com/ver1619/itFrame/compare"
```

---

## What's Inside

| Name | What It Does |
|------|-------------|
| `Comparator[T]` | Interface with a single `Less(a, b T) bool` method |
| `LessFunc[T]` | Turn any function into a `Comparator` |
| `Reverse[T]` | Flips a comparator to sort in the opposite direction |
| `Multi[T]` | Chains multiple comparators for multi-field sorting |
| `Equal` | Checks if two values are equal using a comparator |
| `LessOrEqual` | Checks if `a <= b` using a comparator |

---

## Syntax

### Comparator Interface

```go
type Comparator[T any] interface {
    Less(a, b T) bool
}
```

### LessFunc (quick comparator from a function)

```go
cmp := compare.LessFunc[int](func(a, b int) bool {
    return a < b
})
```

### Reverse

```go
reversed := compare.Reverse[T]{Cmp: originalComparator}
```

### Multi (multi-field sorting)

```go
multi := compare.Multi[T]{
    Comparators: []compare.Comparator[T]{cmp1, cmp2},
}
```

### Utility Functions

```go
compare.Equal[T](cmp, a, b)      // true if a == b under cmp
compare.LessOrEqual[T](cmp, a, b) // true if a <= b under cmp
```

---

## Use Cases

- **Sorting and merging** — `ops.Merge` and `ops.Distinct` need a comparator
- **Multi-field sorting** — sort by last name, then first name, then age
- **Reverse ordering** — flip any comparator for descending order
- **Equality checks** — compare values without writing separate equality logic

---

## Examples

### Example 1 — Basic Integer Comparator

Create a simple comparator for integers and use it to check ordering.

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/compare"
)

func main() {
    cmp := compare.LessFunc[int](func(a, b int) bool {
        return a < b
    })

    fmt.Println(cmp.Less(3, 7))                  // true
    fmt.Println(cmp.Less(7, 3))                  // false
    fmt.Println(compare.Equal(cmp, 5, 5))        // true
    fmt.Println(compare.LessOrEqual(cmp, 3, 3))  // true
    fmt.Println(compare.LessOrEqual(cmp, 5, 3))  // false
}
```

**Output:**
```
true
false
true
true
false
```

---

### Example 2 — Multi-Field Sorting with Reverse

Sort employees by department (ascending), then by salary (descending) using `Multi` and `Reverse`.

```go
package main

import (
    "fmt"
    "github.com/ver1619/itFrame/compare"
)

type Employee struct {
    Name   string
    Dept   string
    Salary int
}

func main() {
    // Sort by department A-Z
    byDept := compare.LessFunc[Employee](func(a, b Employee) bool {
        return a.Dept < b.Dept
    })

    // Sort by salary low-to-high, then reverse for high-to-low
    bySalary := compare.Reverse[Employee]{
        Cmp: compare.LessFunc[Employee](func(a, b Employee) bool {
            return a.Salary < b.Salary
        }),
    }

    // Combine: department first, then salary descending
    combined := compare.Multi[Employee]{
        Comparators: []compare.Comparator[Employee]{byDept, bySalary},
    }

    emps := []Employee{
        {"Alice", "Engineering", 90000},
        {"Bob", "Engineering", 120000},
        {"Charlie", "Marketing", 70000},
    }

    // Check ordering between employees
    fmt.Println(combined.Less(emps[0], emps[1])) // false — same dept, but Bob earns more (reversed = Bob first)
    fmt.Println(combined.Less(emps[1], emps[0])) // true  — Bob before Alice in Engineering (higher salary)
    fmt.Println(combined.Less(emps[0], emps[2])) // true  — Engineering < Marketing
}
```

**Output:**
```
false
true
true
```
