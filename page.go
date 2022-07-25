package fungi

// Page pages stream items according to the given page number (where the first
// page has number 0) and page size.
func Page[T any](pageNumber, pageSize int) StreamIdentity[T] {
	from := pageNumber * pageSize
	to := from + pageSize
	return Range[T](from, to)
}
