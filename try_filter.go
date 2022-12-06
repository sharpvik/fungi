package fungi

// TryFilter stream items using a custom checker function.
// Use it with filtration operations that might error out.
func TryFilter[T any](ok func(T) (bool, error)) StreamIdentity[T] {
	return func(items Stream[T]) Stream[T] {
		return &tryFilter[T]{
			source: items,
			ok:     ok,
		}
	}
}

type tryFilter[T any] struct {
	source Stream[T]
	ok     func(T) (bool, error)
}

func (f *tryFilter[T]) Next() (item T, err error) {
	for {
		item, err = f.source.Next()
		if err != nil {
			return
		}
		var ok bool
		if ok, err = f.ok(item); err != nil || ok {
			return
		}
	}
}
