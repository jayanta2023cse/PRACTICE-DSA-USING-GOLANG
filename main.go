package main

import (
	"fmt"
	"main/programs"
)

func main() {
	// programs.GetNetWork()

	fmt.Println(programs.ReverseString("Token"))
	fmt.Println(programs.ReverseStringOptimized("Kuldeep"))

	fmt.Println(programs.FindDuplicateSol1([]int{1, 3, 4, 2, 2}))
	fmt.Println(programs.FindDuplicateSol1([]int{3, 1, 3, 4, 2}))

}
