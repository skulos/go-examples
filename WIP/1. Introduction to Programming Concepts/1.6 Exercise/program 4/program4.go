// Simple Interest = (Principal(R) x Rate(%) x Time(Y)) / 100
// Interest = (500) * 10% * 6
// Interest = 25
// Total Repayment = 500 + 25 = 525

package main

import (
	"fmt"
)

func main() {

	fmt.Println("================================================================================================")
	fmt.Println("  Simple Interest Calculator")
	fmt.Println("================================================================================================")

	var principal, rate, time int

	fmt.Print("Principal amount: R")
	fmt.Scan(&principal)

	fmt.Print("Interest Rate (%): ")
	fmt.Scan(&rate)

	fmt.Print("Time in years: ")
	fmt.Scan(&time)

	var interest float64 = (float64(principal) * float64(rate) * float64(time)) / 100

	fmt.Println("================================================================================================")
	fmt.Println("Interest is R", interest)

}
