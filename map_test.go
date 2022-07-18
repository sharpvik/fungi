package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	increment := Map(increment)
	stream := increment(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, result)
}

func increment(x int) int {
	return x + 1
}
