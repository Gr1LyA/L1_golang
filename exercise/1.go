package main

import "fmt"

type Human struct {
	name string
}

// Геттер для поля name
func (h *Human) GetName() string {
	return h.name
}

// Сеттер для поля name
func (h *Human) SetName(str string) {
	h.name = str
}

//Встраивание в Action от Human
type Action struct {
	Human
}

func main() {
	var act Action

	act.SetName("Ilya")
	fmt.Println(act.GetName())
}
