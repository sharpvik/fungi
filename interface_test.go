package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream(t *testing.T) {
	stream := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	even := Filter(func(i int) bool { return i%2 == 0 })
	double := Map(func(i int) int { return i * 2 })
	add := Reduce(func(i, j int) int { return i + j }, 0)
	result, err := add(double(even(stream)))
	assert.NoError(t, err)
	assert.Equal(t, 60, result)
}
