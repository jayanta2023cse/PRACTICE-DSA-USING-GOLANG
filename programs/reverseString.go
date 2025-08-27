package programs

// Approach 1:
// we will traverse from the back and add each chracter into a string
func ReverseString(str string) string {
	letters := []rune(str)
	n := len(letters)
	outputStr := ""
	for i := n - 1; i >= 0; i-- {
		outputStr += string(letters[i])
	}
	return outputStr
}

// Approach 2:
// Using two pointers, one is at first index and another is at last index and move them by one place till n/2
func ReverseStringOptimized(s string) string {
	letters := []rune(s)
	n := len(letters)
	for i := 0; i < n/2; i++ {
		letters[i], letters[n-i-1] = letters[n-i-1], letters[i]
	}
	return string(letters)
}
