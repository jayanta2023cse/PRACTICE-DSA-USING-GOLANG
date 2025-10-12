package programs

import "log"

// Approach 1:
// we will traverse from the back and add each chracter into a string

func ReverseString() {
	var str = "Token"

	letters := []rune(str)
	n := len(letters)
	outputStr := ""
	for i := n - 1; i >= 0; i-- {
		outputStr += string(letters[i])
	}

	log.Println("Reversed String is -->", outputStr)
}

// Approach 2:
// Using two pointers, one is at first index and another is at last index and move them by one place till n/2

func ReverseStringOptimized() {
	var s = "Kuldeep"

	letters := []rune(s)
	n := len(letters)
	for i := 0; i < n/2; i++ {
		letters[i], letters[n-i-1] = letters[n-i-1], letters[i]
	}

	log.Println("Reversed String is -->", string(letters))
}
