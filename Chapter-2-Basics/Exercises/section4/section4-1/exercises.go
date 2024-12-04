package exercises

// Exercise 1:
// Write a function ApplyDeliveryFee that takes totalSpending
// as an argument and returns the deliveryFee (either 0 or 250,
//  depending on whether the spending exceeds 500).
func ApplyDeliveryFee(totalSpending int) int {
	if totalSpending < 500 {
		return 250
	} else {
		return 0
	}

}

// Exercise 2:
// Write a function CalculateDiscount that returns a discount based on the totalSpending amount:
// 	- 20% discount for spending over 500,
// 	- 10% for over 200, and
// 	- 5% for less than 200.
func CalculateDiscount(totalSpending int) int {
	if totalSpending > 500 {
		return 20
	} else if totalSpending > 200 {
		return 10
	} else {
		return 5
	}

}

// Exercise 3:
// Write a function CheckEquality that returns "Equal" if two numbers are the same and "Not Equal" otherwise
func CheckEquality(a, b int) string {
	if a == b {
		return "Equal"
	} else {
		return "Not Equal"
	}
}

// Exercise 4:
// Write a function CanBuy that returns true if both balance is greater than or equal to cost and
// age is greater than or equal to 18.
func CanBuy(balance, cost, age int) bool {

	if balance >= cost && age >= 18 {
		return true
	} else {
		return false
	}
}

// Exercise 5:
// Write a function IsEligibleForDiscount that returns true if either membership is "Gold" or totalSpending is greater than 500.
func IsEligibleForDiscount(membership string, totalSpending int) bool {

	if membership == "Gold" || totalSpending > 500 {
		return true
	} else {
		return false
	}
}

// Exercise 6:
// Write a function CheckAge that returns:
// 	- "Minor" if age is less than 18.
// 	- "Adult" if age is between 18 and 50.
//	- "Senior" if age is greater than 60.
func CheckAge(age int) string {
	if age < 18 {
		return "Minor"
	} else if age <= 60 {
		return "Adult"
	} else {
		return "Senior"
	}
}

// Exercise 7:
// Write a function IsValid that returns true if both age is greater than 18 and country is "RSA"
func IsValid(age int, country string) bool {

	if age > 18 && country == "RSA" {
		return true
	} else {
		return false
	}
}

// Exercise 8:
// Write a function IsBiggerThan100 that returns true if the given number is greater than 100
func IsBiggerThan100(number int) bool {
	if number > 100 {
		return true
	} else {
		return false
	}

}

// Exercise 9:
// Write a function IsNotEqual that checks whether two numbers are not equal and returns true
// if they are not equal
func IsNotEqual(a, b int) bool {
	if a != b {
		return true
	} else {
		return false
	}

}

// Exercise 10:
// Write a function WithinRange that checks if a number is within a specified range (inclusive).
func WithinRange(value, min, max int) bool {
	return value >= min && value <= max
}
