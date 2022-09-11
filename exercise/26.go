package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var in string

	for {
		fmt.Scanln(&in)
		start := time.Now()
		fmt.Println("check: ", check(in), time.Now().Sub(start))
		start = time.Now()
		fmt.Println("check1: ", check1(in), time.Now().Sub(start))
	}
}

func check(str string) bool {
	for _, v := range str {
		if strings.Count(str, string(v)) != 1 {
			return false
		}
	}

	return true
}

//храним количество повторений по ключу(rune)
func check1(str string) bool {
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
