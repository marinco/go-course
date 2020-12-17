package hw2

import (
	"errors"
	"fmt"
)

var ToExpensiveErr = errors.New("Too expensive item")

func sendData(numbers chan Shopping) {
	shopList := []Shopping{
		Shopping{"Kruh", 8, 2},
		//Shopping{"Gitara", 150, 1}, // error
		Shopping{"Mlijeko", 15, 1},
	}
	for _, item := range shopList {
		numbers <- item
	}
	close(numbers)
}

func Main() {
	num := make(chan Shopping, 1)
	go sendData(num)
	mostCommon, err := totalCost(100, num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Total cost is ", mostCommon) // rezultat primjera je 31
	}
}

func totalCost(maxPrice int, numbers chan Shopping) (int, error) {
	var sum = 0
	for number := range numbers {
		if maxPrice < number.Price && maxPrice < (number.Price*number.Quantity) {
			return 0, ToExpensiveErr
		}
		sum += number.Price * number.Quantity
		fmt.Println("number is ", number)
	}
	return sum, nil
}
