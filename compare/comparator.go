package compare

type Comparator[T any] interface {
	Less(a, b T) bool
}

type LessFunc[T any] func(a, b T) bool

func (f LessFunc[T]) Less(a, b T) bool {
	return f(a, b)
}

/*
- **Comparator** defines ordering via Less(a,b)
- **LessFunc** lets you use a function directly
*/
