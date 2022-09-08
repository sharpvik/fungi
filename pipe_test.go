package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipe(t *testing.T) {
	var sum int32
	source := SliceStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	filter := Filter(func(i int) bool { return i%2 == 0 })
	mapper := Map(func(i int) int32 { return int32(i) })
	looper := Loop(func(i int32) { sum += i })
	err := Then(Then(
		Pipe(source, filter), mapper), looper).
		Result()
	assert.NoError(t, err)
	assert.Equal(t, int32(sum), sum)
}
