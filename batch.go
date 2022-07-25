package fungi

func Batch[T any](actions ...StreamIdentity[T]) StreamIdentity[T] {
	return func(items Stream[T]) Stream[T] {
		for _, do := range actions {
			items = do(items)
		}
		return items
	}
}
