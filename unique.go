package fungi

func Unique[T comparable](items Stream[T]) Stream[T] {
	return UniqueBy(identity[T])(items)
}

func UniqueBy[T any, K comparable](id func(T) K) StreamIdentity[T] {
	memory := make(map[K]struct{})
	unique := FilterMap(func(item T) (T, bool) {
		key := id(item)
		_, seen := memory[key]
		memory[key] = struct{}{}
		return item, !seen
	})
	return StreamIdentity[T](unique)
}

func identity[T any](item T) T { return item }
