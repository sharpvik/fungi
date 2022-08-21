package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	findEvenNumber := Find(even)
	result, err := findEvenNumber(source)
	assert.NoError(t, err)
	assert.Equal(t, 2, result)
}

func TestFindNoMatch(t *testing.T) {
	source := SliceStream([]int{1, 3, 5, 7, 9})
	findEvenNumber := Find(even)
	_, err := findEvenNumber(source)
	assert.Error(t, err)
}
