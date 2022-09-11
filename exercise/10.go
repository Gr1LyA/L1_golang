package main

import (
	"fmt"
	"math"
)

func main() {
	var arr []float64 = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	res := make(map[int][]float64)

	for _, v := range arr {
		position := int(math.Abs(v / 10)) * 10
		// if _, ok := res[position]; !ok {
		// 	res[position] = make([]float64, 0, 1)
		// }
		res[position] = append(res[position], v)
	}

	fmt.Println(res)
}