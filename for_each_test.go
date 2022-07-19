package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	accumulator := 0
	adder := ForEach(func(i int) { accumulator += i })
	err := adder(source)
	assert.NoError(t, err)
	assert.Equal(t, 55, accumulator)
}
