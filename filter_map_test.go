package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterMap(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	stream := FilterMap(oddPlusOne)(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []int{2, 4, 6, 8, 10}, result)
}

func oddPlusOne(i int) (int, bool) {
	return i + 1, i%2 != 0
}
