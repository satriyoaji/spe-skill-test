package test

// findNeedle finds the index of a needle string in a haystack array of strings.
func FindNeedle(haystack []string, needle string) int {
	for index, str := range haystack {
		if str == needle {
			return index
		}
	}
	return -1
}
