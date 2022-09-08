package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEffect(t *testing.T) {
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	stream := SliceStream(source)
	sum := 0
	sameStream := Effect(func(i int) { sum += i })(stream)
	items, err := CollectSlice(sameStream)
	assert.NoError(t, err)
	assert.Equal(t, source, items)
}
