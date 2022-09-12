package main

import (
	"fmt"
	"math/big"
)
func main() {
	{
		num1 := new(big.Int)
		num2 := new(big.Int)

		num1.SetString("2400000000000000000000000000000000000", 10)
		num2.SetString("2400000000000000000000000000000000000", 10)

		fmt.Println("a =", num1, "b =", num2)

		res := big.NewInt(0)

		fmt.Printf("\n\n")
		fmt.Println("a * b = ", res.Mul(num1, num2))
		fmt.Println("a / b = ", res.Quo(num1, num2))
		fmt.Println("a + b = ", res.Add(num1, num2))
		fmt.Println("a - b = ", res.Sub(num1, num2))
		fmt.Printf("\n\n")
	}

	// Так же если использовать нетипизированные костанты, то в них поместится больше чем в int64.
	// Компилятор будет использовать пакет big. Но они не взаимозаменяемы
	{
		const num1 = 2400000000000000000000000000000000000
		const num2 = 2400000000000000000000000000000000000

		// В случае умножения и суммы Println будет ругаться на переполнение

		// fmt.Println("a * b = ", num1 * num2)
		fmt.Println("a / b = ", num1 / num2)
		// fmt.Println("a + b = ", num1 + num2)
		fmt.Println("a - b = ", num1 - num2)
	}
}