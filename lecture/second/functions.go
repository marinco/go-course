package second

import "fmt"

func IterateSlice(slice []int) {
	for i, i2 := range slice {
		fmt.Println("i,i2", i, i2)
	}
}
