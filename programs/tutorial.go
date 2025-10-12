package programs

import (
	"fmt"
	"slices"
)

// - Given a list of integers, return the highest product of three numbers.

// Example:
// - Input: []int{-10, -10, 1, 3, 2}
//  Output: 300, because -10.-10.3 gives the highest product

func GetHighestProduct() {
	var num = []int{-10, -10, 1, 3, 2}

	n := len(num)
	slices.Sort(num)

	result1 := num[0] * num[1] * num[n-1]
	result2 := num[n-1] * num[n-2] * num[n-3]

	fmt.Printf("Product is %d", Max(result1, result2))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
