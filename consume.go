package fungi

import "io"

// Consume all items of a [Stream], completely ignoring them, until an error is
// received. If the error was [io.EOF], no error will be returned to the caller.
// All other errors are propagated.
func Consume[T any](s Stream[T]) (err error) {
	for err == nil {
		_, err = s.Next()
	}

	if err == io.EOF {
		err = nil
	}

	return
}
