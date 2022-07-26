package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ts := SliceStream(slice)
	rs := SliceStream(slice)
	zip := Zip(add)
	stream := zip(ts)(rs)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, result)
}
