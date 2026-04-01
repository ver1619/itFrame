package core

type Iterator[T any] interface {
	Next() (T, bool)
}

/*DESIGN MODEL:
  - Pull-based: The consumer drives iteration by calling Next().
  - Lazy: Values are produced only when requested.
  - Stateful: Each call to Next() advances internal state.

METHOD CONTRACT:

   Next() (T, bool)

   Returns:
     - (value, true)  → next element is available
     - (zero, false) → iteration is complete (exhausted)

Iteration is single-pass (cannot restart).
No extra memory is used during iteration (efficient for large data).
*/
