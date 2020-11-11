package main

import (
	"fer.hr/go-course/homework/hw1"
	"fer.hr/go-course/lecture/lecture1"
	"fer.hr/go-course/lecture/lecture2"
	"fmt"
	"math"
	"os/user"
)

func main() {

	numbers := [][]int{
		{3, 10, 4, 2},
		{9, 3, 8, 7},
		{15, 14, 13, 12},
	}
	solution := hw1.HappyNumber(numbers)
	fmt.Println(solution)

	shopList := []hw1.Shopping{
		hw1.Shopping{"kruh", 8, 2},
		hw1.Shopping{"mlijeko", 15, 1}}
	solution1, err := hw1.MostExpensive(shopList)
	fmt.Println(solution1) // mlijeko
	fmt.Println(err)       // nil
	solution2 := hw1.TotalCost(shopList)
	fmt.Println(solution2) // 31

	fmt.Println(hw1.MostExpensive(nil))
	fmt.Println(hw1.TotalCost(nil))
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func firstLecture() {
	fmt.Println("My favorite number is ", math.Pi)
	var r float64 = 10
	fmt.Println("Circle area is ", lecture1.CircleArea(r))
	n, s := lecture1.IsPrime(21)
	fmt.Printf("%.f %s\n", n, s)
	n, s = lecture1.IsPrime(23)
	fmt.Printf("%.f %s\n", n, s)
	fmt.Println("Square root from 100 is ", lecture1.Sqrt(100))
	fmt.Println(lecture1.Pow(3, 2, 10))
	fmt.Println(lecture1.Pow(3, 3, 20))
	fmt.Print("Operating system is ")
	lecture1.Os()
	lecture1.ReverseRunes("Hello, world!")

	var m = 6
	fmt.Println(m, "th fibonacci is", lecture1.FibonacciRecursive(m))
	fmt.Println(m, "th fibonacci is", lecture1.FibonacciIterative(m))
	fmt.Println(lecture1.CalculateFibonacci(33))
	fmt.Println(lecture1.CalculateFibonacci(44))
	fmt.Println(lecture1.CalculateFibonacci(55))
	fmt.Println(lecture1.CalculateFibonacci(66))
	fmt.Println(lecture1.Lookup(66))

	fmt.Println(lecture1.FibonacciBinet(33))

	var (
		v1 = lecture1.Point{1, 2}  // tip Tocka
		v2 = lecture1.Point{X: 1}  // Y:0 implicitno
		v3 = lecture1.Point{}      // X:0 i Y:0
		p  = &lecture1.Point{1, 2} // tip *Tocka
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

func secondLecture() {
	var slice = []int{2, 3, 5, 7, 11, 13}

	printSlice(slice)
	slice = slice[:0] // slice kapaciteta 0
	printSlice(slice)
	slice = slice[:4] // proširi slice
	printSlice(slice)
	slice = slice[2:] // izbaci prve dvije vrijednosti
	printSlice(slice)

	var mySlice = make([]int, 0)
	printSlice(mySlice)
	mySlice = append(mySlice, 2, 3, 4)
	printSlice(mySlice)

	lecture2.IterateSlice(mySlice)
	var b, c = 252, 198
	fmt.Println("gcd is", lecture2.Gcd(b, c))

	var m = make(map[string]int)
	m["FER Zagreb"] = 123

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	var d = lecture2.Distance{1, 1, 5, 7}
	fmt.Println(d.AlgDistance())

	var s = lecture2.Student{
		Name:    "Miljenko",
		ECTS:    240,
		JMBAG:   "0036123456",
		Courses: []string{"Matan 1", "Signali i sustavi"},
	}
	fmt.Println(s)
}
