// 5. Calculate the Hypotenuse of a Triangle
// hypotenuse = sqrt(a^2 + b^2)
//		 _
// sqrt /
// tot die mag = ^2

// x ^ 3 = x * x * x
// x ^ 15 = x * x * x * x * x * x* x * x * x * x * x * x* x * x * x

package main

import (
	"fmt"
	"math"
)

func main() {
	// a^2 + b^2 = c^2
	// sqrt(c^2) = c
	// math.Pow(a, 2) + math.Pow(b, 2)

	var a, b int

	fmt.Println("====================================================================")
	fmt.Println("\tThis program will determine the hypotenuse from a and b")
	fmt.Println("====================================================================")

	fmt.Print("\tPlease enter length a: ")
	fmt.Scan(&a)

	fmt.Print("\tPlease enter length b: ")
	fmt.Scan(&b)

	hypotenuse := math.Sqrt(float64(((a * a) + (b * b))))
	// (a * a) + (b * b) = math.Pow(a, 2) + math.Pow(b, 2)

	fmt.Println("\tThe hypotenuse of", a, "and", b, "is", hypotenuse)
}
