package main

import (
	"fmt"
	"sync"
)

// Мьютекс нужен чтоб не было гонки данных при конкурентном доступе к map
// В данном примере использовал RWMutex, он обеспечивает возможность одновременного паралельного чтения
// Однако если идет запись в map, то читать и записывать другие горутины в этот момент не могут
// В этой задаче можно было обойтись и обычным мьютексом, так как чтение из map и нету вовсе.
// В принципе, всегда можно обойтись обычным мьютексом, но тогда параллельно читать не получится.
type MapMutex struct {
	mx sync.RWMutex
	m  map[int]int
}

func NewRWMap() *MapMutex {
	return &MapMutex{
		m: make(map[int]int),
	}
}

func (c *MapMutex) Load(key int) (int, bool) {
	c.mx.RLock()
	val, ok := c.m[key]
	c.mx.RUnlock()
	return val, ok
}

func (c *MapMutex) Store(key int, value int) {
	c.mx.Lock()
	c.m[key] = value
	c.mx.Unlock()
}

func main() {
	fmt.Println("start")

	maprw := NewRWMap()
	wg := &sync.WaitGroup{}

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("routine #", i)
			maprw.Store(1, 1)
		}(i)
	}

	wg.Wait()
}
