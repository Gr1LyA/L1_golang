package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sleep1(time.Second * 2)
	fmt.Println(time.Now().Sub(start))
	sleep2(time.Second * 2)
	fmt.Println(time.Now().Sub(start))
}

func sleep1(n time.Duration) {
	<- time.After(n)
}

func sleep2(n time.Duration) {
	start := time.Now()
	for time.Now().Sub(start) < n {

	}
}