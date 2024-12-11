package testing

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	exercises "github.com/skulos/go-examples/Chapter-2-Basics/Exercises/section4/section4-3/4-3-1"
)

// Helper function to capture printed output
func captureOutput(f func()) string {
	old := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

// Test 1: Print Until N
func TestPrintUntilN(t *testing.T) {
	tests := []struct {
		n              int
		expectedOutput string
	}{
		{5, "1\n2\n3\n4\n5\n"},
		{3, "1\n2\n3\n"},
		{0, "Invalid input\n"},
		{-2, "Invalid input\n"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			output := captureOutput(func() {
				exercises.PrintUntilN(tt.n)
			})
			if output != tt.expectedOutput {
				t.Errorf("Expected:\n%sGot:\n%s", tt.expectedOutput, output)
			}
		})
	}
}

// Test 2: Count Down From N
func TestCountDownFromN(t *testing.T) {
	tests := []struct {
		n              int
		expectedOutput string
	}{
		{3, "3\n2\n1\n"},
		{5, "5\n4\n3\n2\n1\n"},
		{0, "Invalid input\n"},
		{-1, "Invalid input\n"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			output := captureOutput(func() {
				exercises.CountDownFromN(tt.n)
			})
			if output != tt.expectedOutput {
				t.Errorf("Expected:\n%sGot:\n%s", tt.expectedOutput, output)
			}
		})
	}
}

// Test 3: Countdown Timer
func TestCountdownTimer(t *testing.T) {
	tests := []struct {
		n              int
		expectedOutput string
	}{
		{3, "3\n2\n1\nTime's up!\n"},
		{5, "5\n4\n3\n2\n1\nTime's up!\n"},
		{0, ""},
		{-1, ""},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			output := captureOutput(func() {
				exercises.CountdownTimer(tt.n)
			})
			if output != tt.expectedOutput {
				t.Errorf("Expected:\n%sGot:\n%s", tt.expectedOutput, output)
			}
		})
	}
}

// Test 4: Sum Until Limit
func TestSumUntilLimit(t *testing.T) {
	tests := []struct {
		max            int
		expectedOutput string
		expectedSum    int
	}{
		{10, "i: 1 Sum: 1\ni: 2 Sum: 3\ni: 3 Sum: 6\ni: 4 Sum: 10\n", 10},
		{15, "i: 1 Sum: 1\ni: 2 Sum: 3\ni: 3 Sum: 6\ni: 4 Sum: 10\ni: 5 Sum: 15\n", 15},
		{0, "", 0},
		{-5, "", 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("max=%d", tt.max), func(t *testing.T) {
			var output string
			sum := -1

			// Capture output and return sum
			output = captureOutput(func() {
				sum = exercises.SumUntilLimit(tt.max)
			})

			if output != tt.expectedOutput {
				t.Errorf("Expected output:\n%sGot:\n%s", tt.expectedOutput, output)
			}
			if sum != tt.expectedSum {
				t.Errorf("Expected sum: %d, Got: %d", tt.expectedSum, sum)
			}
		})
	}
}
