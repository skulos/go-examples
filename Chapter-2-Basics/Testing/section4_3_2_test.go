package testing

import (
	"fmt"
	"testing"

	exercises "github.com/skulos/go-examples/Chapter-2-Basics/Exercises/section4/section4-3/4-3-2"
)

// Test 1: PrintSomethingNTimes
func TestPrintSomethingNTimes(t *testing.T) {
	tests := []struct {
		n              int
		expectedOutput string
	}{
		{3, "Something\nSomething\nSomething\n"},
		{1, "Something\n"},
		{0, "Invalid input\n"},
		{-2, "Invalid input\n"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			output := captureOutput(func() {
				exercises.PrintSomethingNTimes(tt.n)
			})
			if output != tt.expectedOutput {
				t.Errorf("Expected:\n%sGot:\n%s", tt.expectedOutput, output)
			}
		})
	}
}

// Test 2: PrintMultiplicationTable
func TestPrintMultiplicationTable(t *testing.T) {
	tests := []struct {
		n              int
		expectedOutput string
	}{
		{2, "1 x 1 = 1\n1 x 2 = 2\n2 x 1 = 2\n2 x 2 = 4\n"},
		{3, "1 x 1 = 1\n1 x 2 = 2\n1 x 3 = 3\n2 x 1 = 2\n2 x 2 = 4\n2 x 3 = 6\n3 x 1 = 3\n3 x 2 = 6\n3 x 3 = 9\n"},
		{0, "Invalid input\n"},
		{-1, "Invalid input\n"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			output := captureOutput(func() {
				exercises.PrintMultiplicationTable(tt.n)
			})
			if output != tt.expectedOutput {
				t.Errorf("Expected:\n%sGot:\n%s", tt.expectedOutput, output)
			}
		})
	}
}

// Test 3: CountUntilThreshold
func TestCountUntilThreshold(t *testing.T) {
	tests := []struct {
		n              int
		threshold      int
		expectedOutput string
	}{
		{5, 3, "1\n2\n3\n"},
		{10, 7, "1\n2\n3\n4\n5\n6\n7\n"},
		{0, 3, "Invalid input\n"},
		{5, 0, "Invalid input\n"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d_threshold=%d", tt.n, tt.threshold), func(t *testing.T) {
			output := captureOutput(func() {
				exercises.CountUntilThreshold(tt.n, tt.threshold)
			})
			if output != tt.expectedOutput {
				t.Errorf("Expected:\n%sGot:\n%s", tt.expectedOutput, output)
			}
		})
	}
}

// Test 4: SkipMultiplesOf3
func TestSkipMultiplesOf3(t *testing.T) {
	tests := []struct {
		n              int
		expectedOutput string
	}{
		{5, "1\n2\n4\n5\n"},
		{10, "1\n2\n4\n5\n7\n8\n10\n"},
		{0, "Invalid input\n"},
		{-1, "Invalid input\n"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			output := captureOutput(func() {
				exercises.SkipMultiplesOf3(tt.n)
			})
			if output != tt.expectedOutput {
				t.Errorf("Expected:\n%sGot:\n%s", tt.expectedOutput, output)
			}
		})
	}
}

// Test 5: CountdownWithBlastoff
// Test 5: CountdownWithBlastoff
func TestCountdownWithBlastoff(t *testing.T) {
	tests := []struct {
		n              int
		expectedOutput string
	}{
		{3, "Countdown: 3\nCountdown: 2\nCountdown: 1\nBlastoff!\n"},
		{1, "Countdown: 1\nBlastoff!\n"},
		{0, "Invalid input\n"},
		{-2, "Invalid input\n"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			output := captureOutput(func() {
				exercises.CountdownWithBlastoff(tt.n)
			})

			if output != tt.expectedOutput {
				t.Errorf("got %q, want %q", output, tt.expectedOutput)
			}
		})
	}
}
