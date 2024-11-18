package main

import "fmt"

func main() {

	println("Entry Check")
	println("=========================================================")
	var age int
	var membership bool

	fmt.Print("Enter your age:")
	fmt.Scan(&age)
	println("=========================================================")
	fmt.Print("Do you have a membership (True or False): ")
	fmt.Scan(&membership)

	println("=========================================================")
	if age >= 21 && membership {
		fmt.Printf("Entry granted")
	} else {
		fmt.Println("Entry Denied")
	}

}
