package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := [5]int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)

	f := func(num int) {
		fmt.Println(num * num)
		wg.Done()
	}

	for _, v := range nums {
		wg.Add(1)
		go f(v)
	}

	wg.Wait()
}