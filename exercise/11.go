package main

import "fmt"

func main() {
	arr1 := []any{11, 55, -1, 4, 3, 10, 15, 0, 0, 0, 0, "test"}
	arr2 := []any{-1, 3, 55, 2, -1, -1, "test"}

	fmt.Println(cross(arr1, arr2))
	
}

func cross(arr1, arr2 []any) (result []any) {
	res := make(map[any]int)

	for _, v := range arr1 {
		if res[v] == 0 {
			res[v]++
		}
	}

	for _, v := range arr2 {
		if res[v] == 1 {
			res[v]++
		}
	}

	for k, v := range res {
		if v == 2 {
			result = append(result, k)
		}
	}

	return result
}
