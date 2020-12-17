package lecture1

import (
	"fmt"
	"math"
	"os/user"
)

func Main() {
	fmt.Println("My favorite number is ", math.Pi)
	var r float64 = 10
	fmt.Println("Circle area is ", circleArea(r))
	n, s := isPrime(21)
	fmt.Printf("%.f %s\n", n, s)
	n, s = isPrime(23)
	fmt.Printf("%.f %s\n", n, s)
	fmt.Println("Square root from 100 is ", sqrt(100))
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
	fmt.Print("Operating system is ")
	os()
	reverseRunes("Hello, world!")

	var m = 6
	fmt.Println(m, "th fibonacci is", fibonacciRecursive(m))
	fmt.Println(m, "th fibonacci is", fibonacciIterative(m))
	fmt.Println(calculateFibonacci(33))
	fmt.Println(calculateFibonacci(44))
	fmt.Println(calculateFibonacci(55))
	fmt.Println(calculateFibonacci(66))
	fmt.Println(lookup(66))

	fmt.Println(fibonacciBinet(33))

	var (
		v1 = Point{1, 2}  // tip Tocka
		v2 = Point{X: 1}  // Y:0 implicitno
		v3 = Point{}      // X:0 i Y:0
		p  = &Point{1, 2} // tip *Tocka
	)
	v2.Y = 2
	fmt.Println(v1, p, v2, v3)

	var user, _ = user.Current()
	fmt.Println("user is", user.Name)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	var primes = [6]int{2, 3, 5, 7,
		11, 13}
	fmt.Println(primes)
}
