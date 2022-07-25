package fungi

func Page[T any](pageNumber, pageSize int) StreamIdentity[T] {
	from := pageNumber * pageSize
	to := from + pageSize
	return Range[T](from, to)
}
