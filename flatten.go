package fungi

type flatten[T any] struct {
	buf   []T
	inner Stream[[]T]
}

func (f *flatten[T]) Next() (t T, err error) {
	if len(f.buf) > 0 {
		t = f.buf[0]
		f.buf = f.buf[1:]
		return
	}

	f.buf, err = f.inner.Next()
	if err != nil {
		return
	}

	return f.Next()
}

// Flatten a [Stream] of slices.
func Flatten[T any](s Stream[[]T]) Stream[T] {
	return &flatten[T]{
		inner: s,
	}
}
