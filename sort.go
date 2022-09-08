package fungi

import "sort"

// Sort a stream. Beware, this will attempt to exhaust a given stream, collect
// items into a slice, sort it in place, and stream the resulting slice. There
// is just no other way to do it.
func Sort[T any](less func(a, b T) bool) StreamIdentity[T] {
	return func(items Stream[T]) Stream[T] {
		slice, err := CollectSlice(items)
		if err != nil {
			return ErrorStream[T](err)
		}
		sort.Slice(slice, func(i, j int) bool {
			return less(slice[i], slice[j])
		})
		return SliceStream(slice)
	}
}
