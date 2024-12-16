package fungi

import "io"

type sliced[T any] struct {
	size  int
	inner Stream[T]
}

func (s *sliced[T]) Next() (batch []T, err error) {
	batch, err = CollectSlice(Take[T](s.size)(s.inner))

	if err == nil && len(batch) == 0 {
		err = io.EOF
	}

	return
}

// Slice an underlying [Stream] into a stream of batches of size N.
func Slice[T any](n int) StreamTransformer[T, []T] {
	return func(s Stream[T]) Stream[[]T] {
		return &sliced[T]{
			size:  n,
			inner: s,
		}
	}
}
