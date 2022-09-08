package fungi

// Loop is similar to ForEach but it accepts a function that doesn't return any
// error. It is simply a convenient alias for you to use.
func Loop[T any](handle func(T)) StreamHandler[T] {
	return ForEach(func(item T) error { handle(item); return nil })
}
