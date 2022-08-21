package fungi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	accumulator := 0
	adder := ForEach(func(i int) error { accumulator += i; return nil })
	err := adder(source)
	assert.NoError(t, err)
	assert.Equal(t, 55, accumulator)
}

func TestForEachWithErrorDuringHandling(t *testing.T) {
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	err := ForEach(errorMaker)(source)
	assert.Error(t, err)
}

func errorMaker(i int) error {
	if i == 6 {
		return errors.New(":(")
	}
	return nil
}
