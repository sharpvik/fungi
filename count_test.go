package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	stream := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	count, err := Count(stream)
	assert.NoError(t, err)
	assert.Equal(t, 10, count)
}
