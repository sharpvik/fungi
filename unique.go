package fungi

func Unique[T comparable](items Stream[T]) Stream[T] {
	return UniqueBy(identity[T])(items)
}

func UniqueBy[T any, K comparable](id func(T) K) StreamIdentity[T] {
	memory := make(map[K]struct{})
	return func(items Stream[T]) Stream[T] {
		return FilterMap(func(item T) (T, bool) {
			key := id(item)
			_, seen := memory[key]
			memory[key] = struct{}{}
			return item, !seen
		})(items)
	}
}

func identity[T any](item T) T { return item }
