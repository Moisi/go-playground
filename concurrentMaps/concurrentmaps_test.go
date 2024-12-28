package concurrentmaps

import (
	"testing"
)

const (
	maxConcurrentWriters = 10
	iter                 = 1700
)

func BenchmarkInsert(b *testing.B) {
	// stringsSlice := make([]string,iter) // or slice := make([]int, elems)
	m, doneCh := InitChanMap(maxConcurrentWriters) // this starts a goroutine, user must close channel
	defer close(doneCh)

	for i := 0; i < iter; i++ {
		go func() {
			m.Set("", i)
		}()
	}

}
