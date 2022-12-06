package fungi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryFilterMap(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5})
	stream := TryFilterMap(oddPlusOneUnder5)(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []int{2, 4, 6}, result)
}

func TestTryFilterMapFailure(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	stream := TryFilterMap(oddPlusOneUnder5)(source)
	_, err := CollectSlice(stream)
	assert.Error(t, err)
}

func oddPlusOneUnder5(i int) (int, bool, error) {
	if i > 5 {
		return 0, false, errors.New("cannot work with numbers greater than 5")
	}
	return i + 1, i%2 != 0, nil
}
