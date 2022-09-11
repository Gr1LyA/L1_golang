package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{0, 10, 2, 5, 1, 3, 7}

	sort.Ints(arr)

	fmt.Println(arr)

	fmt.Println(sort.SearchInts(arr, 10))
}