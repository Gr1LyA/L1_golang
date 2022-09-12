package main

import (
	"fmt"
	"strings"
)

func main() {
	in := "snow dog sun"

	//Делим строку по пробелам, получаем массив строк
	words := strings.Split(in, " ")

	length := len(words)

	// Выводим в stdout полученный массив с конца
	for i := range words {
		fmt.Print(words[length-1-i])
		if i != length-1 {
			fmt.Print(" ")
		} else {
			fmt.Print("\n")
		}
	}
}
