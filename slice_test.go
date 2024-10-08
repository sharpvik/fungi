package fungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectSlice(t *testing.T) {
	_, err := CollectSlice[int](nil)
	assert.NoError(t, err)
}
