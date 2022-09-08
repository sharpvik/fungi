package fungi

// Zip two streams using a custom zipper function. Just like similar operations
// in other languages, zipping will stop on error from either one of the source
// streams. Therefore, if one of the source streams is shorter than the other,
// the resulting stream is going to contain as many elements as the shorter
// source stream.
func Zip[T, R, O any](zipper func(T, R) O) func(Stream[T]) StreamTransformer[R, O] {
	return func(itemsT Stream[T]) StreamTransformer[R, O] {
		return func(itemsR Stream[R]) Stream[O] {
			return &zip[T, R, O]{
				ts:     itemsT,
				rs:     itemsR,
				zipper: zipper,
			}
		}
	}
}

type zip[T, R, O any] struct {
	ts     Stream[T]
	rs     Stream[R]
	zipper func(T, R) O
}

func (z *zip[T, R, O]) Next() (item O, err error) {
	t, err := z.ts.Next()
	if err != nil {
		return
	}
	r, err := z.rs.Next()
	if err != nil {
		return
	}
	return z.zipper(t, r), nil
}
