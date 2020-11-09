package first

import (
	"errors"
	"math"
)

var IntegerOverflow = errors.New("Integer overflow")

func CheckInt32(n int) error {
	if n > math.MaxInt32 || n < math.MinInt32 {
		return IntegerOverflow
	}
	return nil
}

func CalculateFibonacci(n int) (string, int) {
	var error = CheckInt32(n)
	if error != nil {
		return error.Error(), 0
	}
	return "Success", FibonacciIterative(n)
}

func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func FibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	var fib = 1
	var previousFib = 1
	for i := 2; i < n; i++ {
		var tmp = fib
		fib += previousFib
		previousFib = tmp
	}
	return fib
}
