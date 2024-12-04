package exercises

// Exercise 1:
// Write a function RewardBasedOnTier that takes a membershipTier
// as an argument ("Basic", "Silver", "Gold", "Platinum")
// and returns the corresponding reward as a string:
// - "Basic": "5% discount voucher"
// - "Silver": "10% discount voucher"
// - "Gold": "20% discount voucher"
// - "Platinum": "30% discount voucher"
// - Any other value should return "No reward".
func RewardBasedOnTier(membershipTier string) string {

	switch membershipTier {
	case "Basic":
		return "5% discount voucher"
	case "Silver":
		return "10% discount voucher"
	case "Gold":
		return "20% discount voucher"
	case "Platinum":
		return "30% discount voucher"
	default:
		return "No reward"
	}
}

// Exercise 2:
// Write a function DescribeDay that takes a day of the week (string)
// and returns a description:
// - "Saturday" or "Sunday": "It's the weekend!"
// - Any other day: "It's a weekday."
func DescribeDay(day string) string {
	var returnStr string
	switch day {
	case "Saturday", "Sunday":
		returnStr = "It's the weekend!"
	default:
		returnStr = "It's a weekday."
	}
	return returnStr
}

// Exercise 3:
// Write a function CategorizeAge that takes an age (int)
// and returns the category:
// - Below 13: "Child"
// - 13 to 19: "Teenager"
// - 20 and above: "Adult"
func CategorizeAge(age int) string {
	switch {
	case age < 13:
		return "Child"
	case age >= 13 && age < 20:
		return "Teenager"
	case age >= 20:
		return "Adult"
	default:
		return ""
	}

}

// Exercise 4:
// Write a function DayType that takes a day (string) and a holiday flag (bool).
// Return "Holiday" if the holiday flag is true,
// "Weekend" for "Saturday" or "Sunday" and holiday is false,
// and "Weekday" otherwise.
func DayType(day string, holiday bool) string {

	return ""
}
