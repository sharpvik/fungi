package fungi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryFilter(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5})
	stream := TryFilter(evenUnder5)(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []int{2, 4}, result)
}

func TestTryFilterFailure(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	stream := TryFilter(evenUnder5)(source)
	_, err := CollectSlice(stream)
	assert.Error(t, err)
}

func evenUnder5(i int) (bool, error) {
	if i > 5 {
		return false, errors.New("cannot work with numbers greater than 5")
	}
	return i%2 == 0, nil
}
