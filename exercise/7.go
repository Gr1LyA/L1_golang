package main

import (
	"sync"
	"time"
	"fmt"
)
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

	wg.Add(1)

	for i := 0; i < 10; i++ {
		go func (i int) {
			wg.Wait()
			fmt.Println("routine #", i)
			maprw.Store(1, 1)
		}(i)
	}

	time.Sleep(time.Second)
	wg.Done()
	time.Sleep(time.Second)
}
