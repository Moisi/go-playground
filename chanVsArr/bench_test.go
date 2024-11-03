package chanvsarr

import (
	"testing"
)

const (
	iter = 1700
)

func BenchmarkResultViaArray(b *testing.B) {
	stringsSlice := make([]string,iter) // or slice := make([]int, elems)
	for i := 0; i < 1000; i++ {
		resultViaArray(i, stringsSlice)
	}
}

func BenchmarkResultViaChan(b *testing.B) {
	stringsCh := make(chan string, iter)
	defer close(stringsCh)
	for i := 0; i < iter ; i++ {
		resultViaChannel(i, stringsCh)
	}

}