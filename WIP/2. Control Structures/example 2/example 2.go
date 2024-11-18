package main

import "fmt"

func main() {
	// var storeIsOpen bool = true

	// // password == "This is my password"
	// for storeIsOpen == true {
	// 	// Execute [while] condition is true
	// }

	// storeIsOpen = false

	var password string = "JBL"
	var userInput string

	fmt.Print("Please enter the password: ")
	fmt.Scan(&userInput)

	for password != userInput {
		fmt.Println("Error - incorrect password")
		fmt.Print("Please enter the password: ")
		fmt.Scan(&userInput)
	}

	fmt.Println("Password correct!")
}
