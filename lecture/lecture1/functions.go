package lecture1

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

func reverseRunes(input string) string {
	r := []rune(input)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
