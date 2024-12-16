package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsume(t *testing.T) {
	var sum int
	stream := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	effect := Effect(func(i int) { sum += i })
	assert.NoError(t, Consume(effect(stream)))
	assert.Equal(t, 55, sum)
}
