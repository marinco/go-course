package lecture2

import "fmt"

func Main() {
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

	iterateSlice(mySlice)
	var b, c = 252, 198
	fmt.Println("gcd is", gcd(b, c))

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

	var d = Distance{1, 1, 5, 7}
	fmt.Println(d.AlgDistance())

	var s = Student{
		Name:    "Miljenko",
		ECTS:    240,
		JMBAG:   "0036123456",
		Courses: []string{"Matan 1", "Signali i sustavi"},
	}
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
