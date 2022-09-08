package fungi

// Count elements of a stream.
func Count[T any](items Stream[T]) (count int, err error) {
	err = Loop(func(_ T) { count++ })(items)
	return
}
