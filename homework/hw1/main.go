package hw1

import "fmt"

func Main() {
	//1
	numbers := []int{1, 2, 3, 4, 5, 1, 2, 3}
	solution := counter(numbers)
	fmt.Println(solution)

	//2
	numbers2 := [][]int{
		[]int{9, 12},
		[]int{3, 4},
	}
	solution2 := happyNumber(numbers2)
	fmt.Println(solution2)

	//3
	shopList := []Shopping{
		Shopping{"kruh", 8, 2},
		Shopping{"mlijeko", 15, 1}}
	solution1, err := mostExpensive(shopList)
	fmt.Println(solution1) // mlijeko
	fmt.Println(err)       // nil
	solution3 := totalCost(shopList)
	fmt.Println(solution3) // 31

}
