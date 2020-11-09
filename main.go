package main

import (
	"fmt"
	"math"

	"fer.hr/go-course/lecture/first"
)

func main() {
	fmt.Println("My favorite number is ", math.Pi)
	var r float64 = 10
	fmt.Println("Circle area is ", first.CircleArea(r))
	n, s := first.IsPrime(21)
	fmt.Printf("%.f %s\n", n, s)
	n, s = first.IsPrime(23)
	fmt.Printf("%.f %s\n", n, s)
	fmt.Println("Square root from 100 is ", first.Sqrt(100))
	fmt.Println(first.Pow(3, 2, 10))
	fmt.Println(first.Pow(3, 3, 20))
	fmt.Print("Operating system is ")
	first.Os()
	first.ReverseRunes("Hello, world!")

	var m = 6
	fmt.Println(m, "th fibonacci is", first.FibonacciRecursive(m))
	fmt.Println(m, "th fibonacci is", first.FibonacciIterative(m))

}
