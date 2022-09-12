package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{10, 11, 12, 13, 14}
	wg := new(sync.WaitGroup)

	// Канал для исходных значений
	ch1 := make(chan int)

	// Канал для квадратов исходных значений
	ch2 := make(chan int)

	done := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			// Получение исходного значения массива
			case x := <-ch1:
				// Отправка удвоенного полученного значения
				ch2 <- (x * 2)
			case <-done:
				fmt.Println("goroutine exit")
				return
			}
		}
	}()

	for _, v := range arr {
		//Отправка исходного значения массива
		ch1 <- v
		// Получение удвоенного значения
		fmt.Println(<-ch2)
	}

	close(done)
	wg.Wait()
}
