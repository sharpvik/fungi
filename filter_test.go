package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	filter := Filter(even)
	stream := filter(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []int{2, 4, 6, 8, 10}, result)
}

func even(i int) bool {
	return i%2 == 0
}
