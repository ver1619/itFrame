package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

func main() {
	p := ops.Peek(core.Range(0, 3, 1))

	v, _ := p.Peek()
	fmt.Println("peek:", v)

	v, _ = p.Next()
	fmt.Println("next:", v)
}
