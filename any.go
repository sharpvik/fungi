package fungi

import "io"

// Any returns true if at least one stream item satisfies given condition.
func Any[T any](pass func(T) bool) StreamReducer[T, bool] {
	return func(items Stream[T]) (ok bool, err error) {
		_, err = Filter(pass)(items).Next()
		ok = err == nil
		if err == io.EOF {
			err = nil
		}
		return
	}
}
