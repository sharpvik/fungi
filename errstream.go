package fungi

type ErrStream[T any] struct {
	err error
}

func (e *ErrStream[T]) Next() (_ T, err error) {
	err = e.err
	return
}
