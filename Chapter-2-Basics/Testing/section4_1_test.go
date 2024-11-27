package testing

import (
	"fmt"
	"testing"

	exercises "github.com/skulos/go-examples/Chapter-2-Basics/Exercises/section4/section4-1"
)

// Test 1: Simple if
func TestApplyDeliveryFee(t *testing.T) {
	tests := []struct {
		totalSpending int
		expectedFee   int
	}{
		{650, 0},
		{439, 250},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Spending=%d", tt.totalSpending), func(t *testing.T) {
			result := exercises.ApplyDeliveryFee(tt.totalSpending)
			if result != tt.expectedFee {
				t.Errorf("Expected %d, got %d", tt.expectedFee, result)
			}
		})
	}
}

// Test 2: Else-if
func TestCalculateDiscount(t *testing.T) {
	tests := []struct {
		totalSpending    int
		expectedDiscount int
	}{
		{600, 20},
		{350, 10},
		{150, 5},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Spending=%d", tt.totalSpending), func(t *testing.T) {
			result := exercises.CalculateDiscount(tt.totalSpending)
			if result != tt.expectedDiscount {
				t.Errorf("Expected %d, got %d", tt.expectedDiscount, result)
			}
		})
	}
}

// Test 3: Check Relational Operator == OR !=
func TestCheckEquality(t *testing.T) {
	tests := []struct {
		a, b           int
		expectedOutput string
	}{
		{5, 5, "Equal"},
		{3, 5, "Not Equal"},
		{10, 10, "Equal"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("a=%d, b=%d", tt.a, tt.b), func(t *testing.T) {
			result := exercises.CheckEquality(tt.a, tt.b)
			if result != tt.expectedOutput {
				t.Errorf("Expected %s, got %s", tt.expectedOutput, result)
			}
		})
	}
}

// Test 4: Using && for Multiple Conditions
func TestCanBuy(t *testing.T) {
	tests := []struct {
		balance, cost, age int
		expectedResult     bool
	}{
		{100, 50, 20, true},
		{100, 50, 17, false},
		{50, 100, 20, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("balance=%d, cost=%d, age=%d", tt.balance, tt.cost, tt.age), func(t *testing.T) {
			result := exercises.CanBuy(tt.balance, tt.cost, tt.age)
			if result != tt.expectedResult {
				t.Errorf("Expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}

// Test 5: Using || for Multiple Conditions
func TestIsEligibleForDiscount(t *testing.T) {
	tests := []struct {
		membership     string
		totalSpending  int
		expectedResult bool
	}{
		{"Gold", 100, true},
		{"Silver", 600, true},
		{"Bronze", 400, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("membership=%s, spending=%d", tt.membership, tt.totalSpending), func(t *testing.T) {
			result := exercises.IsEligibleForDiscount(tt.membership, tt.totalSpending)
			if result != tt.expectedResult {
				t.Errorf("Expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}

// Test 6: Check if Age is in Range (else-if)
func TestCheckAge(t *testing.T) {
	tests := []struct {
		age              int
		expectedCategory string
	}{
		{15, "Minor"},
		{25, "Adult"},
		{65, "Senior"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("age=%d", tt.age), func(t *testing.T) {
			result := exercises.CheckAge(tt.age)
			if result != tt.expectedCategory {
				t.Errorf("Expected %s, got %s", tt.expectedCategory, result)
			}
		})
	}
}

// Test 7: Combining Multiple Conditions with &&
func TestIsValid(t *testing.T) {
	tests := []struct {
		age            int
		country        string
		expectedResult bool
	}{
		{20, "RSA", true},
		{17, "RSA", false},
		{20, "Canada", false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("age=%d, country=%s", tt.age, tt.country), func(t *testing.T) {
			result := exercises.IsValid(tt.age, tt.country)
			if result != tt.expectedResult {
				t.Errorf("Expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}

// Test 8: Returning a Value Based on Comparison (<, >)
func TestIsBiggerThan100(t *testing.T) {
	tests := []struct {
		number         int
		expectedResult bool
	}{
		{150, true},
		{100, false},
		{50, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("number=%d", tt.number), func(t *testing.T) {
			result := exercises.IsBiggerThan100(tt.number)
			if result != tt.expectedResult {
				t.Errorf("Expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}

// Test 9: Using != for Inequality Check
func TestIsNotEqual(t *testing.T) {
	tests := []struct {
		a, b           int
		expectedResult bool
	}{
		{3, 4, true},
		{5, 5, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("a=%d, b=%d", tt.a, tt.b), func(t *testing.T) {
			result := exercises.IsNotEqual(tt.a, tt.b)
			if result != tt.expectedResult {
				t.Errorf("Expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}

// Test 10: Handling Range Checks (>= and <=)
func TestWithinRange(t *testing.T) {
	tests := []struct {
		value, min, max int
		expectedResult  bool
	}{
		{5, 1, 10, true},
		{0, 1, 10, false},
		{15, 1, 10, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("value=%d, min=%d, max=%d", tt.value, tt.min, tt.max), func(t *testing.T) {
			result := exercises.WithinRange(tt.value, tt.min, tt.max)
			if result != tt.expectedResult {
				t.Errorf("Expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}
