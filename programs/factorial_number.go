package programs

import (
	"fmt"
)

func factorialOfANumber(number int) int {
	if number <= 1 {
		return 1
	}
	return number * factorialOfANumber(number-1)
}

func Factorial(n int) {
	a := factorialOfANumber(n)
	fmt.Printf("Factorial of number %d is %d \n", n, a)
}
