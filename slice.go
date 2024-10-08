package fungi

import "io"

// SliceStream transforms any slice into a Stream.
func SliceStream[T any](items []T) Stream[T] {
	return &slice[T]{
		source: items,
	}
}

// CollectSlice exhausts any given Stream, putting items into a slice. At the
// end of this operation, io.EOF is not reported as this function is expected to
// read stream until the end anyways.
func CollectSlice[T any](items Stream[T]) (collected []T, err error) {
	if items == nil {
		return
	}

	var item T
	for err == nil {
		if item, err = items.Next(); err == nil {
			collected = append(collected, item)
		}
	}

	if err == io.EOF {
		err = nil
	}

	return
}

type slice[T any] struct {
	source []T
}

func (s *slice[T]) Next() (item T, err error) {
	if len(s.source) == 0 {
		err = io.EOF
		return
	}
	item = s.source[0]      // get first item
	s.source = s.source[1:] // shrink the slice
	return
}
