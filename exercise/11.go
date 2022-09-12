package main

import "fmt"

func main() {
	arr1 := []any{11, 55, -1, 4, 3, 10, 15, 0, 0, 0, 0, "test"}
	arr2 := []any{-1, 3, 55, 2, -1, -1, "test"}

	fmt.Println(cross(arr1, arr2))
	
}

func cross(arr1, arr2 []any) (result []any) {
	// Инициализируем map
	res := make(map[any]int)

	// Проходимся по массиву и кладем 1 по всем значениям ключей из массива
	for _, v := range arr1 {
		res[v] = 1
	}

	// Проходимся по второму массиву, и если по нужному ключу лежит 1,
	// то есть пересечение, добавляем ключ в итоговое множество
	for _, v := range arr2 {
		if res[v] == 1{
			result = append(result, v)

			// Если убрать строку ниже, то в итоговом множестве могут быть повторения
			res[v]++
		}
	}

	fmt.Println(res)

	return result
}
