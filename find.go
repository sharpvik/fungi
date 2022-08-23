package fungi

func Find[T any](pass func(T) bool) StreamReducer[T, T] {
	return func(items Stream[T]) (match T, err error) {
		return Filter(pass)(items).Next()
	}
}