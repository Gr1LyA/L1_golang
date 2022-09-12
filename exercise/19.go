package main

import "fmt"

func main() {
	var in string

	fmt.Scanln(&in)

	// Конвертируем в []rune, чтоб итерироваться по символам, а не по байтам
	str := []rune(in)

	length := len(str)

	for i := length - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}
	fmt.Print("\n")
}