package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTake(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	take := Take[int](5)
	stream := take(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}
