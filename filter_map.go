package fungi

// FilterMap transforms elements that have been kept by the filtering function.
// In order to do that, it accepts a function that, given a stream item will
//
//  1. evaluate whether it is desired in the output stream
//  2. and if it is, transform it using its internal mapping rules
//
// Here's an example of a simple carry function. It will not allow odd numbers
// to carry into the output stream, meanwhile doubling every even number it
// encounters.
//
//	func doubleEvenNumbers(i int) (mapped int, carry bool) {
//	    carry = i%2 == 0
//	    mapped = i * 2
//	    return
//	}
//
// Let's see it in action:
//
//	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
//	stream := FilterMap(doubleEvenNumbers)(source)
//	result, err := CollectSlice(stream)
//
// We would expect the result to be as follows:
//
//	[]int{4, 8, 12, 16, 10}
//
// Odd numbers from the source stream have been filtered out, while each even
// number has been doubled and carried through.
func FilterMap[I, O any](carry func(I) (O, bool)) StreamTransformer[I, O] {
	filter := Filter(func(item I) (ok bool) {
		_, ok = carry(item)
		return
	})
	mapper := Map(func(item I) (mapped O) {
		mapped, _ = carry(item)
		return
	})
	return func(items Stream[I]) Stream[O] {
		return mapper(filter(items))
	}
}
