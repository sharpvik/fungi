package fungi

func ErrorStream[T any](err error) Stream[T] {
	return &errorStream[T]{err}
}

type errorStream[T any] struct {
	err error
}

func (e *errorStream[T]) Next() (_ T, err error) {
	err = e.err
	return
}
