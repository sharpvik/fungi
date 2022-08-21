package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyTrue(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	hasEvenNumbers := Any(even)
	result, err := hasEvenNumbers(source)
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestAnyFalse(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	justOddNumbers := Filter(odd)
	hasEvenNumbers := Any(even)
	result, err := hasEvenNumbers(justOddNumbers(source))
	assert.NoError(t, err)
	assert.False(t, result)
}

func odd(i int) bool {
	return !even(i)
}
