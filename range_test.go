package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	source := SliceStream(slice)
	stream := Range[int](2, 7)(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, slice[2:7], result)
}
