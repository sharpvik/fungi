package fungi

import "io"

// queue is a clever data structure that allows both item enqueue and dequeue
// operations. It implements the Channel interface.
type queue[T any] struct {
	items []T
}

// Queue is a default constructor for queue.
func Queue[T any]() *queue[T] {
	return QueueOf[T]()
}

// QueueOf allows you to initialise a queue with some elements.
func QueueOf[T any](items ...T) *queue[T] {
	return &queue[T]{
		items: items,
	}
}

// Add some items to the queue.
func (q *queue[T]) Add(items ...T) {
	q.items = append(q.items, items...)
}

// Next item from the queue or error if it's empty.
func (q *queue[T]) Next() (item T, err error) {
	if len(q.items) == 0 {
		err = io.EOF
		return
	}
	item = q.items[0]
	q.items = q.items[1:]
	return
}
