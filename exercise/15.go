package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 20)

	//срежет 100 байт
	// justString = v[:100]

	//если б не было проверки, то при строке чья длина < 100 была бы паника
	if utf8.RuneCountInString(v) >= 100 {
		//срежет 100 символов
		justString = string([]rune(v)[:100])
	}
}

func createHugeString(count int) string {
	return strings.Repeat("ф", count)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
