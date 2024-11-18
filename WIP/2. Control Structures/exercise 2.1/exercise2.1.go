package main

import "fmt"

func main() {
	println("-----------------------------------------------")
	println("Checking for Positive, Negative or number Zero")
	println("-----------------------------------------------")
	var number int

	print("Enter ur desired number: ")
	fmt.Scan(&number)

	println("-----------------------------------------------")

	if number < 0 {
		fmt.Println("Your number is Negative")
	} else if number > 0 {
		fmt.Println("Your number is Positive")
	} else {
		fmt.Println("Your number is Zero")
	}
	println("=================================================================")
}
