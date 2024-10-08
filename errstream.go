package fungi

import "io"

func ErrorStream[T any](err error) Stream[T] {
	return &errorStream[T]{err}
}

func EmptyStream[T any]() Stream[T] {
	return ErrorStream[T](io.EOF)
}

type errorStream[T any] struct {
	err error
}

func (e *errorStream[T]) Next() (_ T, err error) {
	err = e.err
	return
}
