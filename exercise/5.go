package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var n int
	wg := new(sync.WaitGroup)

	// Ввод кол-ва секунд которые программа будет работать
	fmt.Scanln(&n)

	// Канал для передачи и чтения данных
	ch := make(chan int)

	// Канал типа struct{} для прекращения работы горутины ,
	// так как struct{} не занимает память, а передовать нам ничего не надо
	done := make(chan struct{})

	// Функция для получения значений из канала
	f1 := func() {
		defer wg.Done()
		for {
			select {
			// чтение данных из канала
			case x := <-ch:
				fmt.Println(x)

			// В этот кейс функция попадет когда мы закроем канал done, по истечении n секунд
			case <-done:
				fmt.Println("f1 exit")
				return
			}
		}
	}

	// Функция для отправки значений в канал
	f2 := func() {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			// Отправляем данные в канал
			case ch <- i:

			// В этот кейс функция попадет когда мы закроем канал done, по истечении n секунд
			case <-done:
				fmt.Println("f2 exit")
				return
			}
		}
	}

	wg.Add(2)
	// Функция для получения значений из канала
	go f1()

	// Функция для отправки значений в канал
	go f2()

	// Для того, чтоб пройти дальше по телу main,
	// нужно дождаться пока мы сможем прочитать данные из канала <-chan time.Time,
	// возвращаемого функцией: func time.After(d time.Duration) <-chan time.Time,
	// Произойдет это по истечении n секунд

	<-time.After(time.Duration(n) * time.Second)

	// Закрываем канал, для прекращения работы горутин
	close(done)
	wg.Wait()
	fmt.Println("main exit")
}
