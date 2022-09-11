package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string
	var integer int
	var interf interface{}
	var interrf2 any

	//с помощью TypeOf() из пакета reflect
	fmt.Println(reflect.TypeOf(str), reflect.TypeOf(integer), reflect.TypeOf(interf), reflect.TypeOf(interrf2))

	//с помощью switch кострукции которая продемонстрирована в фукции tip()
	fmt.Println(tip(str), tip(integer), tip(interf), tip(interrf2))
}

func tip(a any) string {
	switch a.(type) {
	case int: 
		return "int"
	case string: 
		return "string"
	default: 
	}

	return "<nil>"
}