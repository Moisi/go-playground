package concurrentmaps

import (
	"strconv"
	"testing"
)

const (
	iter = 1700
)

func BenchmarkInsert(b *testing.B) {
	// stringsSlice := make([]string,iter) // or slice := make([]int, elems)
	m := chanMap{}
	for i := 0; i < 1000; i++ {
		m.insert("i:" + strconv.FormatInt(int64(i),10),i)
	}
}
