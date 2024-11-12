package fungi

import "sync"

// Go performs a concurrent [Effect] by spinning up a goroutine for each item.
func Go[T any](effect func(T)) StreamIdentity[T] {
	return Effect(func(item T) { go effect(item) })
}

// GoWait is like [Go] but it additionally accepts a [sync.WaitGroup] to which
// it reports once the effect has completed.
func GoWait[T any](effect func(T), wg *sync.WaitGroup) StreamIdentity[T] {
	return Effect(func(item T) {
		wg.Add(1)
		go func() { defer wg.Done(); effect(item) }()
	})
}
