package main

import (
	"fmt"
)

func main() {
	var name string // "John" of "Cathey"

	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Welcome,", name)
}
