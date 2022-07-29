package fungi

import "io"

// Q is a clever data structure that implements Stream and allows item enqueue
// operation.
type Q[T any] struct {
	items []T
}

// Queue is a default constructor for Q.
func Queue[T any]() *Q[T] {
	return QueueOf[T]()
}

// QueueOf allows you to initialise a queue with some elements.
func QueueOf[T any](items ...T) *Q[T] {
	return &Q[T]{
		items: items,
	}
}

// Add some items to the queue.
func (q *Q[T]) Add(items ...T) {
	q.items = append(q.items, items...)
}

// Next item from the queue or error if it's empty.
func (q *Q[T]) Next() (item T, err error) {
	if len(q.items) == 0 {
		err = io.EOF
		return
	}
	item = q.items[0]
	q.items = q.items[1:]
	return
}
