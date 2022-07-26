package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	source := SliceStream([]int{5, 2, 3, 1, 4})
	sort := Sort(func(a, b int) bool { return a < b })
	stream := sort(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}
