package fungi

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGo(t *testing.T) {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	total := 0
	slice := []int{1, 2, 3, 4, 5}
	stream := SliceStream(slice)

	wg.Add(5)
	count := Go(func(i int) {
		defer wg.Done()
		mutex.Lock()
		defer mutex.Unlock()
		total += i
	})

	consume := Loop(Nop[int])
	require.NoError(t, consume(count(stream)))
	wg.Wait()
	require.Equal(t, 15, total)
}

func TestGoWait(t *testing.T) {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	total := 0
	slice := []int{1, 2, 3, 4, 5}
	stream := SliceStream(slice)

	count := GoWait(func(i int) {
		mutex.Lock()
		defer mutex.Unlock()
		total += i
	}, &wg)

	consume := Loop(Nop[int])
	require.NoError(t, consume(count(stream)))
	wg.Wait()
	require.Equal(t, 15, total)
}
