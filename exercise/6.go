package main

import (
	"fmt"
	"time"
	"sync"
	"sync/atomic"
)

func main() {
	case1()
	case2()
}


func case1() {
	done := make(chan struct{})
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <- done:
				fmt.Println("goroutine exit")
				return 
			}
		}
	}()
	
	time.Sleep(time.Second * 2)

	close(done)
	wg.Wait()
}

func case2() {
	stop := int64(0)

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if atomic.LoadInt64(&stop) == 1 {
				fmt.Println("goroutine exit")
				return 
			}
		}
	}()

	time.Sleep(time.Second * 2)

	atomic.AddInt64(&stop, 1)
	wg.Wait()
}