package fungi

import "io"

func Concat[T any](streams ...Stream[T]) Stream[T] {
	return &concat[T]{
		streams: streams,
	}
}

type concat[T any] struct {
	streams []Stream[T]
}

func (c *concat[T]) Next() (item T, err error) {
	if len(c.streams) == 0 {
		err = io.EOF
		return
	}
	if item, err = c.streams[0].Next(); err != io.EOF {
		return
	}
	c.streams = c.streams[1:]
	return c.Next()
}
