package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	reduce := Reduce(add, 0)
	result, err := reduce(source)
	assert.NoError(t, err)
	assert.Equal(t, 55, result)
}

func add(a, b int) int {
	return a + b
}
