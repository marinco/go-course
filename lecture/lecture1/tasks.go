package lecture1

import (
	"errors"
	"math"
)

var IntegerOverflow = errors.New("Integer overflow")

const (
	Fib33 = 3524578
	Fib44 = 701408733
	Fib55 = 139583862445
	Fib66 = 27777890035288
)

func checkInt32(n int) error {
	if n > math.MaxInt32 || n < math.MinInt32 {
		return IntegerOverflow
	}
	return nil
}

func calculateFibonacci(n int) (string, int) {
	var error = checkInt32(n)
	if error != nil {
		return error.Error(), 0
	}
	return "Success", fibonacciIterative(n)
}

func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func fibonacciIterative(n int) int {
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

func fibonacciBinet(n float64) float64 {
	return (math.Pow((1+math.Sqrt(5))/2, n) - math.Pow((1-math.Sqrt(5))/2, n)) / math.Sqrt(5)
}

func lookup(n int) int {
	switch n {
	case 33:
		return Fib33
	case 44:
		return Fib44
	case 55:
		return Fib55
	case 66:
		return Fib66
	default:
		return fibonacciRecursive(n)
	}
}
