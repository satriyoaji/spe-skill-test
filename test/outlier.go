package test

// findOutlier finds the outlier in an array of integers.
func FindOutlier(arr []int) interface{} {
	// Check if the array has fewer than 3 elements.
	if len(arr) < 3 {
		return false
	}

	// Initialize counters for odd and even numbers.
	oddCount := 0
	evenCount := 0

	// Initialize variables to store the last odd and even numbers seen.
	var lastOdd, lastEven int

	// Iterate through the array and count odd and even numbers.
	for _, num := range arr {
		if num%2 == 0 {
			evenCount++
			lastEven = num
		} else {
			oddCount++
			lastOdd = num
		}
	}

	// Determine the outlier based on the counts.
	if evenCount == 1 {
		return lastEven
	}
	if oddCount == 1 {
		return lastOdd
	}

	// If no outlier exists, return false.
	return false
}
