package lecture2

import "fmt"

func iterateSlice(slice []int) {
	for i, i2 := range slice {
		fmt.Println("i,i2", i, i2)
	}
}
