package fungi

// Stream allows you to receive items one by one. Once it runs out of items,
// error returned will be io.EOF. Depending on the underlying implementation, it
// might return other errors during a call to Next (e.g. connection errors in
// case this stream depends on a remote server for item fetching).
//
// If you have a custom type that you wish to use with this library, you only
// have to write a converter function that looks sort of like this:
//
//	func Streamline(from MyCustomType) fungi.Stream[T]
//
// then, you will be able to use all fungi operations on the resulting Stream:
//
//	var stream fungi.Stream[int] = Streamline(new(MyCustomType))
//	even        := fungi.Filter(func(i int) bool { return i%2 == 0 })
//	double      := fungi.Map(func(i int) int { return i * 2 })
//	add         := fungi.Reduce(func(i, j int) int { return i + j }, 0)
//	result, err := add(double(even(stream)))
type Stream[T any] interface {
	Next() (T, error)
}

// Here are some function types that transform streams.
type (
	StreamTransformer[I, O any] func(Stream[I]) Stream[O]
	StreamIdentity[T any]       StreamTransformer[T, T]
	StreamReducer[I, O any]     func(Stream[I]) (O, error)
	StreamHandler[T any]        func(Stream[T]) error
)
