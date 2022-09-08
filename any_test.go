package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyTrue(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	result, err := Any(even)(source)
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestAnyFalse(t *testing.T) {
	source := SliceStream([]int{1, 3, 5, 7, 9})
	result, err := Any(even)(source)
	assert.NoError(t, err)
	assert.False(t, result)
}

func odd(i int) bool {
	return !even(i)
}
