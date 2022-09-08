package fungi

type Indexed[T any] struct {
	Index int
	Item  T
}

func Index[T any]() func(T) *Indexed[T] {
	index := -1
	return func(item T) *Indexed[T] {
		index++
		return &Indexed[T]{
			Index: index,
			Item:  item,
		}
	}
}

// Enumerate turns any stream of items into a stream of Indexed items.
func Enumerate[T any](items Stream[T]) Stream[*Indexed[T]] {
	return Map(Index[T]())(items)
}

// Deindex unpacks an item from its Indexed box.
func Deindex[T any](indexed *Indexed[T]) T {
	return indexed.Item
}

// Denumerate removes indexing information from a stream of Indexed items.
func Denumerate[T any](items Stream[*Indexed[T]]) Stream[T] {
	return Map(Deindex[T])(items)
}
