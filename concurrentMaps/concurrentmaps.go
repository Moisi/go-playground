package concurrentmaps

import (
	"fmt"
	"reflect"
)

type elem struct {
	key string
	val int
}

type concurentMapInterface interface {
	Get(key string)
	Set(key string, val int)
	Delete(key string)
}

// impl with channels no mutex:
type ChanMap struct {
	postKey chan elem
	close   chan interface{}

	// underlying non concurrent map:
	m map[string]int
}

func InitChanMap(maxConcurrentWriters int) (ChanMap, chan interface{}) {
	m := ChanMap{}
	// user would use this channel to stop the map's bg process
	m.close = make(chan interface{})

	// concurrent writers will post SET data to this channel
	// which will be read sequntially by manager and stored.
	// more counters would block waiting for channel read.
	m.postKey = make(chan elem, maxConcurrentWriters)

	go func() {
		for {
			select {
			case <-m.close:
				break
			case p := <-m.postKey:
				fmt.Printf("Received a %q /n", reflect.TypeOf(p).Name())
			}
		}
	}()

	return m, m.close
}

func (m *ChanMap) Set(key string, val int) {
	fmt.Println("key:", key, ". val:", val)
	if len(m.postKey) == cap(m.postKey) {
		fmt.Println("channel full, this call would block:", key, ":", val)
		// TODO(metric): increase block counter?
	}
	m.postKey <- elem{key, val}
	return
}
