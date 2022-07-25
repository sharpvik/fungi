package fungi

import "io"

// ForEach uses the custom handler function given to it to process every item
// from a stream.
func ForEach[T any](handle func(T) error) StreamHandler[T] {
	return func(items Stream[T]) (err error) {
		var item T
		for err == nil {
			if item, err = items.Next(); err == nil {
				err = handle(item)
			}
		}
		if err == io.EOF {
			err = nil
		}
		return
	}
}
