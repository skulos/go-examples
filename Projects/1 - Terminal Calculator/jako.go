package main

import (
	"fmt"
	"math"
)

func Calculator() {

	println("Calculator Ready")
	println("===========================================")
	var operator string
	fmt.Println("Choose opoperator +,-,*,/,%")
	fmt.Scan(&operator)

	var num1, num2 float64
	fmt.Println("Enter two numbers")
	fmt.Scan(&num1)
	fmt.Print(operator)
	fmt.Scan(&num2)
	var Refresh bool
	// small letters for variable names

	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	case "%":
		fmt.Println(math.Mod(num1, num2))
	default:
		fmt.Println("Invalid Operator Entry")

	}
	fmt.Println("Refresh (true/false)")
	fmt.Scan(&Refresh)

}
