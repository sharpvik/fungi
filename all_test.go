package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllTrue(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	justOddNumbers := Filter(odd)
	allNumbersAreOdd := All(odd)
	result, err := allNumbersAreOdd(justOddNumbers(source))
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestAllFalse(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	allNumbersAreEven := All(even)
	result, err := allNumbersAreEven(source)
	assert.NoError(t, err)
	assert.False(t, result)
}
