package concurrentmaps

import "fmt"

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
type chanMap struct {
	postKey chan elem

	// underlying non concurrent map:
	m map[string]int
}

func (m *chanMap) insert(key string, val int) {
	fmt.Println("key:", key, ". val:", val)
	return
}
