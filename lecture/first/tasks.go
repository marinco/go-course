package first

func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func FibonacciIterative(n int) int {
	if n <= 1 {
		return 1
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
