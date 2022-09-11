package main

import (
	"fmt"
	"strings"
)

func main() {
	in := "snow dog sun"

	words := strings.Split(in, " ")

	var res string

	length := len(words)

	for i := range words {
		res += words[length - 1 - i]
		if i != length - 1 {
			res += " "
		}
	}

	fmt.Print(res)
}