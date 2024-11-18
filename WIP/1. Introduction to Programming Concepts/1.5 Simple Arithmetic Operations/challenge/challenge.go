/* kg floats

1. Sell bananas @14.99kg
2. Sell apples @7.50kg
3. Plastic Bag @0.15 per kg

Calculate the total cost?

==============================================================
Bananas    Apples
Xkg     +   Ykg     = Zkg

How much plastic bags to add to total? (Z x 0.15)
==============================================================
*/

package main

import "fmt"

func main() {
	var x float64 = 14.99
	var y float64 = 7.50
	var z float64 = 1.50
	var amountB float64
	var amountA float64
	var amountP float64

	fmt.Println("============================================================================")
	fmt.Println("   Welcome to my digital till")
	fmt.Println("============================================================================")

	fmt.Println("Bananas per kg R", x)
	fmt.Print("Enter kg: ")
	fmt.Scan(&amountB)
	fmt.Print("\bkg\n")
	fmt.Println("Total: R", amountB*x, "\n")

	fmt.Println("-----------------------------------------------------------------------------")

	fmt.Println("Apples per kg R", y)
	fmt.Scan(&amountA)
	fmt.Println("R", amountA*y)

	fmt.Println("============================================================================")

	// fmt.Println("plastic bag x1 R", z)
	// fmt.Scan(&amountP)
	// fmt.Println("R", amountP*z)

	// Zkg  = Xkg + Ykg
	amountP = amountB + amountA
	fmt.Println("Total kg ", amountP, " @ R1.50 per kg")
	fmt.Println("Plastic bags added R", z*amountP)

	fmt.Print("Total R", (amountA*y)+(amountB*x)+(amountP*z))
}
