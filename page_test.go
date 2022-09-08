package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPage(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//                               ^^^^^^^
	//                          page #2 of size 3
	source := SliceStream(slice)
	stream := Page[int](2, 3)(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []int{7, 8, 9}, result)
}
