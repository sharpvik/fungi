package fungi

// Map transforms every item from a stream using a custom mapper function.
func Map[I, O any](transform func(I) O) StreamTransformer[I, O] {
	return func(items Stream[I]) Stream[O] {
		return &mapper[I, O]{
			source:    items,
			transform: transform,
		}
	}
}

type mapper[I, O any] struct {
	source    Stream[I]
	transform func(I) O
}

func (m *mapper[I, O]) Next() (item O, err error) {
	origin, err := m.source.Next()
	if err != nil {
		return
	}
	return m.transform(origin), nil
}
