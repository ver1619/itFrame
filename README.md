# *itFrame*

itFrame is a pull-based iterator framework in Go designed for:
- lazy evaluation
- composable iteration
- minimal allocation overhead

---

## Current Version: v0.1.0

### Features
- Generic Iterator interface
- SliceIterator
- RangeIterator (end-exclusive)

---


### Requirements

```go
go 1.23+
```


### Installation

```go
go get github.com/ver1619/itFrame
```


----

### To run tests

```go
go test ./...
```

---


### Examples


#### 1. Slice Iterator

```go
package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
)

func main() {
	it := core.NewSliceIterator([]int{10, 20, 30})

	for {
		val, ok := it.Next()
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
```

#### 2. Range Iterator

```go
package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
)

func main() {
	it := core.NewRangeIterator(0, 5, 1)

	for {
		val, ok := it.Next()
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
```


Run:
```go
go run <file.go>
```


