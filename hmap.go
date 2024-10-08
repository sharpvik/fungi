package fungi

import "io"

// CollectMap exhaustively collects items from a given stream into a builtin map
// type. As its first argument, it requires a function that knows how to take an
// item and split it into map Key and Value of appropriate type.
func CollectMap[
	T, V any,
	K comparable,
](keyval func(T) (K, V)) StreamReducer[T, map[K]V] {
	return func(items Stream[T]) (hmap map[K]V, err error) {
		hmap = make(map[K]V)

		if items == nil {
			return
		}

		var item T
		for err == nil {
			if item, err = items.Next(); err == nil {
				key, value := keyval(item)
				hmap[key] = value
			}
		}

		if err == io.EOF {
			err = nil
		}

		return
	}
}
