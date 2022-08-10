package fungi

import "io"

// Any returns true if at least one stream item satisfies given condition.
func Any[T any](pass func(T) bool) StreamReducer[T, bool] {
	return func(items Stream[T]) (ok bool, err error) {
		check := ForEach(func(item T) error {
			if ok = pass(item); ok {
				return io.EOF
			}
			return nil
		})
		err = check(items)
		return
	}
}
