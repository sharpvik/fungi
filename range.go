package fungi

// Range picks out stream items between specified boundary indexes, similar to
// the way builtin slicing works in go:
//
//	slice := []int{1, 2, 3, 4, 5}
//	slice_ := slice[2:3]
func Range[T any](from, to int) StreamIdentity[T] {
	drop := Drop[T](from)
	take := Take[T](to - from)
	return Batch(drop, take)
}
