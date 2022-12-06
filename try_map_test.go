package fungi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryMap(t *testing.T) {
	source := SliceStream([]int{2, 4, 6, 8, 10})
	stream := TryMap(incrementEven)(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []int{3, 5, 7, 9, 11}, result)
}

func TestTryMapFailure(t *testing.T) {
	source := SliceStream([]int{2, 4, 5, 6, 8, 10})
	stream := TryMap(incrementEven)(source)
	_, err := CollectSlice(stream)
	assert.Error(t, err)
}

func incrementEven(x int) (int, error) {
	if x%2 != 0 {
		return 0, errors.New("cannot increment odd numbers")
	}
	return x + 1, nil
}
