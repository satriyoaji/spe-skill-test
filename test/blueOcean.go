package test

// blueOcean removes all values from the first array that are present in the second array.
func BlueOcean(arr1 []int, arr2 []int) []int {
	// Create a map to store elements from the second array.
	excludeMap := make(map[int]bool)

	// Populate the map with elements from the second array.
	for _, num := range arr2 {
		excludeMap[num] = true
	}

	// Initialize an empty result array.
	var result []int

	// Iterate through the first array and add elements to the result that are not in the map.
	for _, num := range arr1 {
		if !excludeMap[num] {
			result = append(result, num)
		}
	}

	return result
}
