package testing

import (
	"fmt"
	"testing"

	exercises "github.com/skulos/go-examples/Chapter-2-Basics/Exercises/section4/section4-2"
)

// Test 1: Membership Rewards
func TestRewardBasedOnTier(t *testing.T) {
	tests := []struct {
		membershipTier string
		expectedReward string
	}{
		{"Basic", "5% discount voucher"},
		{"Silver", "10% discount voucher"},
		{"Gold", "20% discount voucher"},
		{"Platinum", "30% discount voucher"},
		{"None", "No reward"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Tier=%s", tt.membershipTier), func(t *testing.T) {
			result := exercises.RewardBasedOnTier(tt.membershipTier)
			if result != tt.expectedReward {
				t.Errorf("Expected %s, got %s", tt.expectedReward, result)
			}
		})
	}
}

// Test 2: Describe Day
func TestDescribeDay(t *testing.T) {
	tests := []struct {
		day          string
		expectedDesc string
	}{
		{"Saturday", "It's the weekend!"},
		{"Sunday", "It's the weekend!"},
		{"Monday", "It's a weekday."},
		{"Friday", "It's a weekday."},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Day=%s", tt.day), func(t *testing.T) {
			result := exercises.DescribeDay(tt.day)
			if result != tt.expectedDesc {
				t.Errorf("Expected %s, got %s", tt.expectedDesc, result)
			}
		})
	}
}

// Test 3: Age Category
func TestCategorizeAge(t *testing.T) {
	tests := []struct {
		age         int
		expectedCat string
	}{
		{10, "Child"},
		{15, "Teenager"},
		{20, "Adult"},
		{40, "Adult"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Age=%d", tt.age), func(t *testing.T) {
			result := exercises.CategorizeAge(tt.age)
			if result != tt.expectedCat {
				t.Errorf("Expected %s, got %s", tt.expectedCat, result)
			}
		})
	}
}

// Test 4: Day Type
func TestDayType(t *testing.T) {
	tests := []struct {
		day          string
		holiday      bool
		expectedType string
	}{
		{"Monday", false, "Weekday"},
		{"Saturday", false, "Weekend"},
		{"Saturday", true, "Holiday"},
		{"Wednesday", true, "Holiday"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Day=%s,Holiday=%v", tt.day, tt.holiday), func(t *testing.T) {
			result := exercises.DayType(tt.day, tt.holiday)
			if result != tt.expectedType {
				t.Errorf("Expected %s, got %s", tt.expectedType, result)
			}
		})
	}
}
