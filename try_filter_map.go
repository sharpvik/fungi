package fungi

// TryFilterMap works just like FilterMap, however, it accepts a mapping
// function that might error out, in which case, calls to Next on this stream
// will return a non-nil error.
func TryFilterMap[I, O any](carry func(I) (O, bool, error)) StreamTransformer[I, O] {
	return func(items Stream[I]) Stream[O] {
		return &tryFilterMap[I, O]{
			source: items,
			carry:  carry,
		}
	}
}

type tryFilterMap[I, O any] struct {
	source Stream[I]
	carry  func(I) (O, bool, error)
}

func (t *tryFilterMap[I, O]) Next() (item O, err error) {
	for {
		var origin I
		if origin, err = t.source.Next(); err != nil {
			return
		}
		if item, ok, err := t.carry(origin); err != nil || ok {
			return item, err
		}
	}
}
