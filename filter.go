package fungi

// Filter stream items using a custom checker function.
func Filter[T any](ok func(T) bool) StreamIdentity[T] {
	return func(items Stream[T]) Stream[T] {
		return &filter[T]{
			source: items,
			ok:     ok,
		}
	}
}

type filter[T any] struct {
	source Stream[T]
	ok     func(T) bool
}

func (f *filter[T]) Next() (item T, err error) {
	for {
		if item, err = f.source.Next(); err != nil || f.ok(item) {
			return
		}
	}
}
