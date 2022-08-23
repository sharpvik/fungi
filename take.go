package fungi

import "io"

// Take just first N items of the Stream. The resulting stream is going to start
// returning io.EOF once N items have been depleted through calls to Next.
func Take[T any](n int) StreamIdentity[T] {
	return func(items Stream[T]) Stream[T] {
		return &take[T]{
			source: items,
			number: n,
		}
	}
}

type take[T any] struct {
	source Stream[T]
	number int
}

func (t *take[T]) Next() (item T, err error) {
	if t.number == 0 {
		err = io.EOF
		return
	}
	if item, err = t.source.Next(); err == nil {
		t.number--
	}
	return
}
