package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string
	var integer int
	var interf interface{}

	//с помощью TypeOf() из пакета reflect
	fmt.Println(reflect.TypeOf(str), reflect.TypeOf(integer), reflect.TypeOf(interf))

	//с помощью switch кострукции которая продемонстрирована в фукции tip()
	fmt.Println(tip(str), tip(integer), tip(interf))
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