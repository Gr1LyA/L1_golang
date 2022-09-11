package main

import "fmt"

//если порядок не важен
func remove1(s []int, i int) []int {
	length := len(s)

	s[i] = s[length - 1]

	return s[:length - 1]
}

//если важен
func remove2(s []int, i int) []int {
	return append(s[:i], s[i + 1:]...)
}

func main() {
	arr := []int{1, 2, 3, 4 ,5, 6}

	fmt.Println(remove2(arr, 2))

	fmt.Println(remove1(arr, 2))
}