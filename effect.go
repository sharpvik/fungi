package fungi

func Effect[T any](effect func(T)) StreamIdentity[T] {
	return StreamIdentity[T](Map(func(item T) T { effect(item); return item }))
}
