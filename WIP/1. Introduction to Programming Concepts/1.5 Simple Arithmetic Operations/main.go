package main

import "fmt"

func main() {
	x := 10 // 1 2 3 4 5
	y := 5  // -2 -1 0 1 2 3

	var z float64 = float64(y) / float64(x)

	// 100 int
	// 100 float64

	// 100 int
	// 100.0 float64

	fmt.Println("Addition:", x+y)       // 10 + 5 = 15
	fmt.Println("Subtraction:", x-y)    // 10 - 5 = 5
	fmt.Println("Reverse:", y-x)        // 5 - 10 = -5
	fmt.Println("Multiplication:", x*y) // 10 x 5 = 50
	fmt.Println("Division:", x/y)       // 10 / 5 = 2
	fmt.Println("Reverse:", z)          // 5 / 10 = 0.5 (1/2) ~ 0
}
