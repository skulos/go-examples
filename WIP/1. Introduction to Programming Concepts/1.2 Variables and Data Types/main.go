package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var variable1 int = -100 // -2 -1 0 1 2 3 4 5

	var variable2 string = "I love New York" // "abc" ""

	var variable3 float64 = 3.14 // 0 3.14

	var variable4 bool = true // true false

	var variable5 uint = 123145 // 0 1 2 3 4 ...

	fmt.Println(variable1)
	fmt.Println(variable4)
	fmt.Println(variable3)
	fmt.Println(variable2)
	fmt.Println(variable5)

	// var newVariable int = 100
	// newVariable := 100

	// var router GinGonicServer.Router = ...
	// router := ...

	var a, b, c int

	a, b, c = 1, 1, 1
	// a = 1
	// b = 1
	// c = 1

	//========================================================
	var randomInt int = rand.Int()

	fmt.Println(randomInt)

	var isEven bool = randomInt.Even() // false

}
