package main

import (
	"main/programs"
)

func main() {
	// programs.GetNetWork()

	// Reverse a string
	programs.ReverseString("Token")
	programs.ReverseStringOptimized("Kuldeep")

	// Find Duplicate element
	programs.FindDuplicateSol1([]int{1, 3, 4, 2, 2})
	programs.FindDuplicateSol1([]int{3, 1, 3, 4, 2})

	// Race Condition
	programs.RaceConditionWithoutMutes()
	programs.RaceConditionWithMutes()

	// Usage of Select Statement in Go
	programs.SelectStatmentExample()

}
