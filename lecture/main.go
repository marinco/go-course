package main

import (
	"fmt"
	"math"
	"runtime"
)

func circleArea(r float64) float64 {
	return math.Pow(r, 2) * math.Pi
}

func isPrime(n float64) (float64, string) {
	var i float64
	for i = 2; i < math.Floor(math.Sqrt(n)); i++ {
		if (math.Mod(n, i)) == 0 {
			return i, "Number is not prime"
		}
	}
	return 0, "Number is prime"
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	//v not visible here
	return lim
}

func os() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("os is ", os)

	}
}

func testDefer() {
	fmt.Println("Begin")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("End")
}

func main() {
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
	fmt.Print("Operyting system is ")
	os()
	testDefer()
}
