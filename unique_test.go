package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	source := SliceStream([]int{1, 1, 2, 2, 3, 3})
	stream := Unique(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, result)
}
