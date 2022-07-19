package fungi

import "io"

type StreamHandler[T any] func(Stream[T]) error

// ForEach uses the custom handler function given to it to process every item
// from a stream.
func ForEach[T any](handle func(T)) StreamHandler[T] {
	return func(items Stream[T]) (err error) {
		var item T
		for err == nil {
			if item, err = items.Next(); err == nil {
				handle(item)
			}
		}
		if err == io.EOF {
			err = nil
		}
		return
	}
}
