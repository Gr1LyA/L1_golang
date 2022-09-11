package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	nums := [5]int{2, 4, 6, 8, 10}

	fmt.Println(way1(nums))

	fmt.Println(way2(nums))
}

//с помощью атомарных операций
func way1(nums [5]int) (res int64) {
	wg := new(sync.WaitGroup)

	f := func(num int) {
		atomic.AddInt64(&(res), int64(num * num))
		wg.Done()
	}

	for _, v := range nums {
		wg.Add(1)
		go f(v)
	}

	wg.Wait()
	return
}

//с помощью мьютексов
func way2(nums [5]int) (res int64) {
	wg := new(sync.WaitGroup)
	var mx sync.Mutex

	f := func(num int) {
		mx.Lock()
		res += int64(num * num)
		mx.Unlock()
		wg.Done()
	}

	for _, v := range nums {
		wg.Add(1)
		go f(v)
	}

	wg.Wait()
	return
}