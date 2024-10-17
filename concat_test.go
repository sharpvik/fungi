package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	first := SliceStream(source[:5])
	second := SliceStream(source[5:])
	stream := Concat(first, second)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, source, result)
}

func TestConcatWithEmptyStream(t *testing.T) {
	source := []int{1, 2, 3}
	empty := EmptyStream[int]()
	stream := Concat(empty, SliceStream(source))
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, source, result)
}
