package testing

import (
	"fmt"
	"testing"

	exercises "github.com/skulos/go-examples/Chapter-2-Basics/Exercises/section4/section4-1"
)

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
