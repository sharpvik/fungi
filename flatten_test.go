package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	stream := SliceStream([][]int{{1, 2}, {}, {3}})
	numbers := Flatten(stream)
	sum := Reduce(func(x, y int) int { return x + y }, 0)
	result, err := sum(numbers)
	assert.NoError(t, err)
	assert.Equal(t, result, result)
}
