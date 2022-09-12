package main

import (
	"fmt"
	"sort"
)

type Human struct {
	Name	string
	Age		int
}

func main() {
	{
		arr := []int{0, 10, 2, 5, 1, 3, 7}

		sort.Ints(arr)

		fmt.Println(arr)
	}

	{
		arr := []Human{
			{Name:"ilya", Age: 66},
			{Name:"olya", Age: 1},
			{Name:"andrey", Age: 55},
			{Name:"evkakiy", Age: 3},
			{Name:"Yuliy", Age: 0},
		}

		// Для сортировки slice нужно передать функцию less
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].Age < arr[j].Age
		})

		fmt.Println(arr)
	}
}