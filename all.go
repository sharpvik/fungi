package fungi

// All returns true if and only if all stream items satisfy given condition.
func All[T any](pass func(T) bool) StreamReducer[T, bool] {
	return func(item Stream[T]) (ok bool, err error) {
		fail, err := Any(func(item T) bool { return !pass(item) })(item)
		ok = !fail
		return
	}
}
