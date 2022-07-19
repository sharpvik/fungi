package fungi

type Indexed[T any] struct {
	Index int
	Item  T
}

// Deindex unpacks an item from its Indexed box.
func Deindex[T any](indexed *Indexed[T]) T {
	return indexed.Item
}

// Enumerate turns any stream of items into a stream of Indexed items.
func Enumerate[T any](items Stream[T]) Stream[*Indexed[T]] {
	index := 0
	return Map(func(item T) (indexed *Indexed[T]) {
		indexed = &Indexed[T]{
			Index: index,
			Item:  item,
		}
		index++
		return
	})(items)
}

// Denumerate removes indexing information from a stream of Indexed items.
func Denumerate[T any](items Stream[*Indexed[T]]) Stream[T] {
	return Map(Deindex[T])(items)
}
