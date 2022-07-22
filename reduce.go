package fungi

import "io"

// Reduce a stream by combining items in a special way.
func Reduce[I, O any](combine func(O, I) O, def O) StreamReducer[I, O] {
	return func(items Stream[I]) (O, error) {
		return (&reduce[I, O]{
			source:      items,
			accumulator: def,
			combine:     combine,
		}).all()
	}
}

type reduce[I, O any] struct {
	source      Stream[I]
	accumulator O
	combine     func(O, I) O
}

func (r *reduce[I, O]) all() (result O, err error) {
	var item I
	for err == nil {
		if item, err = r.source.Next(); err == nil {
			r.accumulator = r.combine(r.accumulator, item)
		}
	}
	if err == io.EOF {
		return r.accumulator, nil
	}
	return
}
