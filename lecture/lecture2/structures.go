package lecture2

import (
	"fmt"
	"math"
)

type Step struct {
	B int
	A int
	R int
}

type Point struct {
	Lat  float64
	Long float64
}

type Distance struct {
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
}

func (d Distance) AlgDistance() float64 {
	return math.Sqrt(math.Pow(d.X1-d.X2, 2) + math.Pow(d.Y1-d.Y2, 2))
}

type Student struct {
	Name    string
	ECTS    int
	JMBAG   string
	Courses []string
}

func (s Student) String() string {
	return fmt.Sprintf("Student %v (%v), number of ECTS: %v, applied following courses: %v", s.Name, s.JMBAG, s.ECTS, s.Courses)
}
