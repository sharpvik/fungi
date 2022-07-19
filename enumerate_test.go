package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumerate(t *testing.T) {
	source := Enumerate(SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	filter := Filter(oddIndex)
	stream := Denumerate(filter(source))
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []int{2, 4, 6, 8, 10}, result)
}

func oddIndex(indexed *Indexed[int]) bool {
	return indexed.Index%2 != 0
}
