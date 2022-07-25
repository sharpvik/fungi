package fungi

func Range[T any](from, to int) StreamIdentity[T] {
	drop := Drop[T](from)
	take := Take[T](to - from)
	return func(items Stream[T]) Stream[T] {
		return take(drop(items))
	}
}
