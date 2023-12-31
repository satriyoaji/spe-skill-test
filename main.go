package main

import (
	"fmt"
	"spe-skill-test/test"
)

func main() {
	// Narcissistic test
	fmt.Println(test.Narcissistic(153))  // Should print true
	fmt.Println(test.Narcissistic(111))  // Should print false
	fmt.Println(test.Narcissistic(1634)) // Should print true

	// Find Outlier test
	fmt.Println(test.FindOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36})) // Should return 11
	fmt.Println(test.FindOutlier([]int{160, 3, 1719, 19, 11, 13, -21})) // Should return 160
	fmt.Println(test.FindOutlier([]int{11, 13, 15, 19, 9, 13, -21}))    // Should return false

	// find needle in haystack
	fmt.Println(test.FindNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "blue"))  // Should return 1
	fmt.Println(test.FindNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "green")) // Should return -1
	fmt.Println(test.FindNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "red"))   // Should return 0

	// Blue Ocean reverse
	fmt.Println(test.BlueOcean([]int{1, 2, 3, 4, 6, 10}, []int{1}))     // Should return [2, 3, 4, 6, 10]
	fmt.Println(test.BlueOcean([]int{1, 5, 5, 5, 5, 3}, []int{5}))      // Should return [1, 3]
	fmt.Println(test.BlueOcean([]int{1, 2, 3, 4, 6, 10}, []int{1, 10})) // Should return [2, 3, 4, 6]
}
