package programs

import "fmt"

func findFibonacciNumber(number int64) int64 {
	if number <= 1 {
		return number
	}
	return findFibonacciNumber(number-1) + findFibonacciNumber(number-2)
}

func Fibonacci(n int64) {
	result := findFibonacciNumber(n)

	fmt.Printf("%dth fibonacci number is %d \n", n, result)
}
