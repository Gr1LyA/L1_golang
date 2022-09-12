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

	// ввод кол-ва воркеров
	fmt.Scanln(&n)

	// Канал для передачи и чтения данных
	ch := make(chan string)

	// Канал типа struct{} для прекращения работы горутины ,
	// так как struct{} не занимает память, а передовать нам ничего не надо
	done := make(chan struct{})

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			var out string

			for {
				select {
				// чтение данных из канала
				case out = <- ch :
					fmt.Println(out)

				// В этот кейс функция попадет когда мы закроем канал done, после получения Ctrl+C
				case <- done:
					fmt.Println("goroutine exit")
					return
				}
			}
		}()
	}


	// Создаем канал для получение сигнала
	signalChan := make(chan os.Signal)

	// Канал будет получать уведомление только о Ctrl+C
	signal.Notify(signalChan, os.Interrupt)

	for {
		select{
		// Если был отправлен Interrupt, то закрываем done для завершения горутин.
		// Заверешение горутин ожидаем с помощью wg.Wait()
		case <- signalChan:
			fmt.Println("main exit")
			close(done)
			wg.Wait()
			return

		// Отправляем данные в канал
		case ch <- "test ":
		}
	}
}
