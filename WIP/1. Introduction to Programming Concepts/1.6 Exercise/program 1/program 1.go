// 6. Determine Even or Odd Number

// 1 - Odd
// 2 - Even
// 3 - Odd
// 4 - Even

// If multiple of 2, then even. (As deelbaar deur 2 - geen brekkdeel)

// 2 * X (1, 2, 3, 4, 5, ...)

// Y % 2 = 0 OR 1 (0 - Even, 1 - Odd)

package main

import "fmt"

func main() {

	var number int

	fmt.Println("====================================================================")
	fmt.Println("\tThis program will determine if an integer is even or odd")
	fmt.Println("====================================================================")

	fmt.Print("\tPlease enter an integer: ") //number here
	fmt.Scan(&number)

	var remainder int = number % 2 // 0 or 1 (modulus)

	// % 8
	// 0 1 2 3 4 5 6 7

	// if remainder == 0 {
	// 	fmt.Println("\nThe number", number, "is even!")
	// }

	// if remainder == 1 {
	// 	fmt.Println("\nThe number", number, "is odd!")
	// }

	if remainder == 0 {
		fmt.Println("\nThe number", number, "is even!")
	} else {
		fmt.Println("\nThe number", number, "is odd!")
	}

}
