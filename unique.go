package fungi

// Unique returns a stream of unique items. The incoming stream must be a stream
// of comparable items.
func Unique[T comparable](items Stream[T]) Stream[T] {
	return UniqueBy(identity[T])(items)
}

// UniqueBy accepts a stream of any items and returns a stream of unique items.
// To be able to do that efficiently, it requires a function that can convert
// a stream item into something comparable. It's like a hashing function, but it
// doesn't have to be.
//
// For example, let's say there's a struct called User:
//
//	type User struct {
//	    ID int // unique database identifier
//	    Name string
//	}
//
// You can create a stream of unique users like so:
//
//	var usersStream fungi.Stream[*User] = service.GetUsers()
//	//  usersStream might contain duplicates
//	uniqueUsers := fungi.UniqueBy(func(u *User) int { return u.ID })
//	uniqueUsersStream := uniqueUsers(usersStream)
func UniqueBy[T any, K comparable](id func(T) K) StreamIdentity[T] {
	memory := make(map[K]struct{})
	unique := FilterMap(func(item T) (T, bool) {
		key := id(item)
		_, seen := memory[key]
		memory[key] = struct{}{}
		return item, !seen
	})
	return StreamIdentity[T](unique)
}

func identity[T any](item T) T { return item }
