package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForward(t *testing.T) {
	stream := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	filter := Filter(even)
	queuer := Queue[int]()
	err := Forward[int](queuer)(filter(stream))
	assert.NoError(t, err)
	result, err := CollectSlice[int](queuer)
	assert.NoError(t, err)
	assert.Equal(t, []int{2, 4, 6, 8, 10}, result)
}
