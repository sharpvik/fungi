package fungi

// Find an element that satisfies given boolean condition. Here is how it
// returns values based on the search outcome.
//
//  1. Item found      -> (item = item,        err = nil)
//  2. Item not found  -> (item = { garbage }, err = io.EOF)
//  3. Error in stream -> (item = { garbage }, err = error)
func Find[T any](pass func(T) bool) StreamReducer[T, T] {
	return func(items Stream[T]) (match T, err error) {
		return Filter(pass)(items).Next()
	}
}
