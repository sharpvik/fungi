package fungi

// Drop N items and return resulting Stream.
func Drop[T any](n int) StreamIdentity[T] {
	return func(items Stream[T]) Stream[T] {
		for i := 0; i < n; i++ {
			if _, err := items.Next(); err != nil {
				return &ErrStream[T]{err}
			}
		}
		return items
	}
}
