package main

import "fmt"

func main() {
	var a, b int = 1, 0

	fmt.Println(a, b)

	a, b = b, a
	fmt.Println(a, b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}
