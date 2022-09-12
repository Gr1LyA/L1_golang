package main

import (
	"fmt"
	"math"
)

// Инкапсуляция в Go реализуется на уровне пакета.
// То что начинается со строчного символа будет инкапсулированно,
// Обращаться к ним можно будет только внутри пакета
type Point struct {
	x, y float64
}

func New(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	p1 := New(0, 0)
	p2 := New(5, 5)

	fmt.Println(Distance(*p1, *p2))
}
