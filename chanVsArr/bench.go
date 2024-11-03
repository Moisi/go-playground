package chanvsarr

import (
	"fmt"
	"strconv"
	"sync"
)

func task(idx int, result *string) {
	tmp := strconv.Itoa(idx)
	result = &tmp
}

func resultViaArray(con int, ret []string) {
	var wg sync.WaitGroup
	// multi writer:
	wg.Add(con)

	for i := 0; i < con; i++ {
		go func() {
			task(i, &ret[i])
			wg.Done()
		}()
	}

	wg.Wait()

	// single reader - by order
	for i := 0; i < con; i++ {
		save := ret[i]
		if len(save) > 4 {
			fmt.Println(save)
		}
	}
}

func taskCh(idx int, ch chan string) {
	tmp := strconv.Itoa(idx)
	ch <- tmp
}

func resultViaChannel(con int, ret chan string) {
	var wg sync.WaitGroup
	wg.Add(con)
	// multi writer:
	for i := 0; i < con; i++ {
		go func() {
			taskCh(i, ret)
			wg.Done()
		}()
	}
	wg.Wait()

	// single unordered reader:
	for i := 0; i < con; i++ {
		save := <-ret
		if len(save) > 4 {
			fmt.Println(save)
		}
	}
}
