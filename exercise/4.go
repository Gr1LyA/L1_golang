package main

import (
	"fmt"
	"sync"
	"os/signal"
	"os"
)

func main() {
	var n int
	var wg sync.WaitGroup

	fmt.Scanln(&n)

	ch := make(chan string)
	done := make(chan struct{})

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			var out string

			for {
				select {
				case out = <- ch :
					fmt.Println(out)
				case <- done:
					fmt.Println("goroutine exit")
					return
				}
			}
		}()
	}


	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)

	for {
		select{
		case <- signalChan:
			fmt.Println("main exit")
			close(done)
			wg.Wait()
			return
		case ch <- "test ":
		}
	}
}
