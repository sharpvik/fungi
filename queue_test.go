package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	queue := Queue[int]()
	for _, i := range slice {
		queue.Add(i)
	}
	odds := Filter(func(i int) bool { return i%2 != 0 })
	sum := Reduce(func(i, j int) int { return i + j }, 0)
	result, err := sum(odds(queue))
	assert.NoError(t, err)
	assert.Equal(t, 25, result)
}
