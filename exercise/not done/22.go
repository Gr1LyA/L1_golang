package main

import (
	"fmt"
	"math/big"
)
//так же если использовать нетипизированные костанты то в них поместится больше чем в int64. Компилятор будет использовать пакет big
func main() {
	{
		num1 := big.NewInt(2 << 22)
		num2 := big.NewInt(2 << 22 + 1)

		fmt.Println("num1=", num1, "num2=", num2)

		res := big.NewInt(0)

		fmt.Printf("\n\n")
		fmt.Println(res.Mul(num1, num2))
		fmt.Println(res.Quo(num1, num2))
		fmt.Println(res.Add(num1, num2))
		fmt.Println(res.Sub(num1, num2))
		fmt.Printf("\n\n")
	}

	{
		//до 1 << 64 - 1
		num1 := int64(2 << 22)
		num2 := int64(2 << 22 + 1)

		fmt.Println(num1 * num2)
		fmt.Println(num1 / num2)
		fmt.Println(num1 + num2)
		fmt.Println(num1 - num2)
	}
}