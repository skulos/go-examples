package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	// Variables
	var choice int
	var num1, num2, result float64
	var repeatChoice string
	var repeat bool = true
	var invalidChoice = false

	// Menu - Screen 1
	for repeat {
		Clear()
		choice = 0
		num1, num2, result = 0.0, 0.0, 0.0
		repeatChoice = ""

		fmt.Println("===========================================")
		fmt.Println(" Project 1: Calculator")
		fmt.Println("===========================================")

		fmt.Println(" Select a operation:")
		fmt.Println("\t1. +")
		fmt.Println("\t2. -")
		fmt.Println("\t3. x")
		fmt.Println("\t4. /")
		fmt.Println("\t5. % (modulus)")
		// History
		fmt.Println("\t6. Exit Program\n")

		if invalidChoice == true {
			fmt.Println(" Error: Invalid Choice")
			invalidChoice = false
		}

		fmt.Print(" Choice: ")
		fmt.Scan(&choice)

		if choice == 6 {
			return // exit
		}

		// Screen 2
		Clear()

		fmt.Println("===========================================")
		fmt.Println(" Project 1: Calculator")
		fmt.Println("===========================================")

		switch choice {
		case 1:
			fmt.Println("  __  +  __\n\n")
		case 2:
			fmt.Println("  __  -  __\n\n")
		case 3:
			fmt.Println("  __  x  __\n\n")
		case 4:
			fmt.Println("  __  /  __\n\n")
		case 5:
			fmt.Println("  __  %  __\n\n")
		// History case
		// Choice for repeat or not
		default:
			invalidChoice = true
			continue
		}

		fmt.Print("Enter the first number: ")
		fmt.Scan(&num1)

		// Screen 3
		Clear()

		fmt.Println("===========================================")
		fmt.Println(" Project 1: Calculator")
		fmt.Println("===========================================")

		fmt.Print(" ", num1, " ")
		switch choice {
		case 1:
			fmt.Println("+  __\n\n")
		case 2:
			fmt.Println("-  __\n\n")
		case 3:
			fmt.Println("x  __\n\n")
		case 4:
			fmt.Println("/  __\n\n")
		case 5:
			fmt.Println("%  __\n\n")
		}

		fmt.Print("Enter the second number: ")
		fmt.Scan(&num2)

		// Screen 4
		var invalid bool = true
		for invalid {
			Clear()

			fmt.Println("===========================================")
			fmt.Println(" Project 1: Calculator")
			fmt.Println("===========================================")

			if num2 == 0 && choice == 4 { // X / 0
				fmt.Println("Cannot divide by zero!")
			} else if num2 == 0 && choice == 5 {
				fmt.Println("Cannot use modulus with zero!")
			} else {

				fmt.Print(" ", num1, "  ")

				switch choice {
				case 1:
					fmt.Print("+")
					// choice, num1, num2
					result = num1 + num2
					//  num1 +
				case 2:
					fmt.Print("-")
					result = num1 - num2
				case 3:
					fmt.Print("x")
					result = num1 * num2
				case 4:
					fmt.Print("/")
					result = num1 / num2
				case 5:
					fmt.Print("%")
					result = math.Mod(num1, num2)
				}

				fmt.Print("  ", num2, "  =  ", result)
				//  num2 = result
			}

			fmt.Print("\n\n Do you want to do another calculation? (Y/N)  ")
			fmt.Scan(&repeatChoice)

			switch repeatChoice {
			case "Y", "y", "Yes", "yes":
				// if repeatChoice == "Y" || repeatChoice == "y" || repeatChoice == "yes" || repeatChoice == "Yes"
				// repeat
				repeat = true
				invalid = false
			case "N", "n", "No", "no":
				// exit
				repeat = false
				invalid = false
				// exit = true
			default:
				// invalid input
				invalid = true
			}
		} // end of for invalid loop
	} // end of for repeat loop
}
