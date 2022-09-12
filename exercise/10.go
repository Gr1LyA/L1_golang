package main

import (
	"fmt"
)

func main() {
	var arr []float64 = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Инициализируем map, в которой будут хранится сгрупированные значения исходного массива
	res := make(map[int][]float64)

	for _, v := range arr {
		// Определяем подможество, которому соответствует значение
		position := int(v) / 10 * 10

		// Добавляем значение в нужный слайс
		res[position] = append(res[position], v)
	}

	fmt.Println(res)
}