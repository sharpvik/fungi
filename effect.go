package fungi

// Effect applies a function to each item of the stream without altering or
// consuming that stream, such that future stream processing is possible.
func Effect[T any](effect func(T)) StreamIdentity[T] {
	return StreamIdentity[T](Map(func(item T) T { effect(item); return item }))
}
