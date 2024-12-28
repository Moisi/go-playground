package concurrentmaps

import (
	"sync"
	"testing"
)

const (
	maxConcurrentWriters = 2
	iter                 = 20
)

func BenchmarkInsertSameKeyMulti(b *testing.B) {
	// stringsSlice := make([]string,iter) // or slice := make([]int, elems)
	m, doneCh := InitChanMap(maxConcurrentWriters) // this starts a goroutine, user must close channel
	defer close(doneCh)

	wg := sync.WaitGroup{}
	for i := 0; i < iter; i++ {
		wg.Add(2)
		go func() {
			m.Set("", i)
			wg.Done()
		}()

		go func() {
			m.Set(m.Get(""))
			wg.Done()
		}()
	}

	// should finish all map tasks before killing it, other wise
	// if there are blocking writers they would block forever
	wg.Wait()
}
