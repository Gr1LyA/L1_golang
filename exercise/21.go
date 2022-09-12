package main

import "fmt"

// Интерфейс который нам нужен
type inter interface {
	SayHello()
}

// Структура, к которой нужен адаптер
type s1 struct {
}

func (s *s1) Say(str string) {
	fmt.Println(str)
}

// Адаптер для структуры s1
type adapter struct {
	*s1
}

// Метод, реализующий нужный нам интерфейс
func (s *adapter) SayHello() {
	s.Say("Hello")
}

func main() {
	var s s1

	s.Say("Hello")

	adapt := adapter{&s}
	smth(&adapt)
}

func smth(i inter) {
	i.SayHello()
}
