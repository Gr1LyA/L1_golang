package main

import "fmt"

func main() {
	strs := []string {"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(getUniqArr(strs))
}

// Возвращет массив с уникальными элементами
func getUniqArr(strs []string) (res []string) {
	// Инициализируем map, так как значения нас не интересует оно будет типа struct{}
	m := make(map[string]struct{})

	// Оставляем значения по ключам
	for _, v := range strs {
		m[v] = struct{}{}
	}

	// Так как ключи уникальные в результате цикла получаем массив с уникальными элементами
	for k := range m {
		res = append(res, k)
	}

	return
}