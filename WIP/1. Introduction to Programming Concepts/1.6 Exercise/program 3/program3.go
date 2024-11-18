package main

import "fmt"

func main() {
	var a int //Lenth//
	var b int //Width//

	println("------------------------------------------------------------------------")
	println("Calculating the Perimeter of a Rectangle")
	println("------------------------------------------------------------------------")
	println("2 x (length + width)")
	println("------------------------------------------------------------------------")

	print("Enter Lenth: ")
	fmt.Scan(&a)

	print("Enter Width: ")
	fmt.Scan(&b)

	println("------------------------------------------------------------------------")

	print("The total Perimeter of the triangle is: ")
	println(2 * (a + b))
}
