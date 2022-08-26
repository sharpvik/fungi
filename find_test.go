package fungi

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	result, err := Find(even)(source)
	assert.NoError(t, err)
	assert.Equal(t, 2, result)
}

func TestFindNoMatch(t *testing.T) {
	source := SliceStream([]int{1, 3, 5, 7, 9})
	_, err := Find(even)(source)
	assert.Equal(t, io.EOF, err)
}
