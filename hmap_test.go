package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectMap(t *testing.T) {
	source := SliceStream([]int{1, 2, 3})
	result, err := CollectMap(duplicate)(source)
	assert.NoError(t, err)
	assert.Equal(t, map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}, result)
}

func duplicate(i int) (int, int) {
	return i, i
}
