package fungi

func Range[T any](from, to int) StreamIdentity[T] {
	drop := Drop[T](from)
	take := Take[T](to - from)
	return Batch(drop, take)
}
