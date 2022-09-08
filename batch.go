package fungi

// Batch allows you to sequentially apply StreamIdentity functions to a stream
// in order to receive the final StreamIdentity function.
func Batch[T any](actions ...StreamIdentity[T]) StreamIdentity[T] {
	return func(items Stream[T]) Stream[T] {
		for _, do := range actions {
			items = do(items)
		}
		return items
	}
}
