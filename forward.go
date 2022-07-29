package fungi

// Forward items from a Stream into a given Collector of the same type.
func Forward[T any](into Collector[T]) StreamHandler[T] {
	return Loop(func(item T) { into.Add(item) })
}
