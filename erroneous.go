package fungi

type erroneous[T any] struct {
	err error
}

func (e *erroneous[T]) Next() (_ T, err error) {
	err = e.err
	return
}
