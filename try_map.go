package fungi

// TryMap transforms every item from a stream using a custom mapper function.
func TryMap[I, O any](try func(I) (O, error)) StreamTransformer[I, O] {
	return func(items Stream[I]) Stream[O] {
		return &tryMapper[I, O]{
			source: items,
			try:    try,
		}
	}
}

type tryMapper[I, O any] struct {
	source Stream[I]
	try    func(I) (O, error)
}

func (m *tryMapper[I, O]) Next() (item O, err error) {
	origin, err := m.source.Next()
	if err != nil {
		return
	}
	return m.try(origin)
}
