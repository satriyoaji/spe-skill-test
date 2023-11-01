package test

import (
	"math"
	"strconv"
)

// Narcissistic checks if a number is narcissistic or not.
func Narcissistic(num int) bool {
	// Convert the integer to a string to find its length (number of digits).
	numStr := strconv.Itoa(num)
	numDigits := len(numStr)

	// variable to hold the sum.
	sum := 0

	for _, digitRune := range numStr {
		// Convert the rune to an integer.
		digitInt, _ := strconv.Atoi(string(digitRune))

		// Add the digit raised to the power of numDigits to the sum.
		sum += int(math.Pow(float64(digitInt), float64(numDigits)))
	}

	// Check if the sum is equal to the original number.
	return sum == num
}
