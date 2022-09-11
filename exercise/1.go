package main

import "fmt"

type Human struct {
	name string
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) SetName(str string) {
	h.name = str
}


type Action struct {
	Human
}

func main() {
	var act Action

	act.SetName("Ilya")
	fmt.Println(act.GetName())
}