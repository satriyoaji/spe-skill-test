package main

import (
	"fmt"
	"spe-skill-test/test"
)

func main() {
	fmt.Println(test.Narcissistic(153))  // Should print true
	fmt.Println(test.Narcissistic(111))  // Should print false
	fmt.Println(test.Narcissistic(1634)) // Should print true
}
