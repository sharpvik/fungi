package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	stream := SliceStream([]int{1, 2, 3})
	batched := Slice[int](2)(stream)

	batch, err := batched.Next()
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2}, batch)

	batch, err = batched.Next()
	assert.NoError(t, err)
	assert.Equal(t, []int{3}, batch)

	_, err = batched.Next()
	assert.Error(t, err)
}
