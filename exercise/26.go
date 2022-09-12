package main

import (
	"fmt"
	"strings"
)

func main() {
	var in string

	for {
		fmt.Scanln(&in)
		fmt.Println("check: ", check(in))

		fmt.Println("check1: ", check1(in))
	}
}

// Воспользовался функцией strings.Count,
// которая возвращает количество повторений подстроки в строке
func check(str string) bool {

	// Возвращает строку c нижним регистром
	str = strings.ToLower(str)

	for _, v := range str {
		if strings.Count(str, string(v)) != 1 {
			return false
		}
	}

	return true
}

//храним количество повторений по ключу(rune)
func check1(str string) bool {

	// Возвращает строку c нижним регистром
	str = strings.ToLower(str)

	m := make(map[rune]int)

	for _, v := range str {
		if m[v] == 1 {
			return false
		} else {
			m[v]++
		}
	}

	return true
}
