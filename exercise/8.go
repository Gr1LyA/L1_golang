package main

import "fmt"

func main() {
	var n int64 = 1
	i := 3

	fmt.Printf("%b\n", n)
	//установить бит в единичку
	n |= 1 << i

	fmt.Printf("%b\n", n)
	//установить бит в ноль
	n &= ^(1 << i)

	fmt.Printf("%b\n", n)
}