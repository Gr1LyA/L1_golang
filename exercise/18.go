package main

import (
	"fmt"
	"sync"
)

type integer struct {
	x int
	sync.Mutex
}

func (i *integer) Inkrement() {
	i.Lock()
	i.x++
	i.Unlock()
}

func main() {
	var num integer
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 50; i++ {
				num.Inkrement()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(num.x)

}