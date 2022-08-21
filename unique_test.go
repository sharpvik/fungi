package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	source := SliceStream([]int{1, 1, 2, 2, 3, 3})
	stream := Unique(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestUniqueBy(t *testing.T) {
	users := []*user{
		{0, "Maria"},
		{1, "Maria"},
		{0, "Lexie"},
		{1, "Lexie"},
	}
	source := SliceStream(users)
	uniqueByID := UniqueBy(func(u *user) int { return u.id })
	stream := uniqueByID(source)
	result, err := CollectSlice(stream)
	assert.NoError(t, err)
	assert.Equal(t, []*user{users[0], users[1]}, result)
}

type user struct {
	id   int
	name string
}
