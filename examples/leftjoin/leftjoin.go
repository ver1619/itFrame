package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

type User struct {
	ID int
}

type Order struct {
	UserID int
}

func main() {
	users := core.Slice([]User{{1}, {2}})
	orders := core.Slice([]Order{{1}, {1}})

	it := ops.LeftJoin(
		users,
		orders,
		func(u User) int { return u.ID },
		func(o Order) int { return o.UserID },
	)

	for {
		v, ok := it.Next()
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
