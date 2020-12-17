package hw1

import (
	"errors"
	"math"
)

func counter(numbers []int) map[int]int {
	var pairs = make(map[int]int)
	for _, element := range numbers {
		pairs[element]++
	}
	return pairs
}

func happyNumber(matrica [][]int) int {
	for i := 0; i < len(matrica); i++ {
		var min = math.MaxInt32
		for j := 0; j < len(matrica[i]); j++ {
			//find min element in row
			if matrica[i][j] <= min {
				min = matrica[i][j]
			}
		}

		for j := 0; j < len(matrica[i]); j++ {
			var max = math.MinInt32
			for k := 0; k < len(matrica); k++ {
				//find max element in column
				if matrica[k][j] >= max {
					max = matrica[k][j]
				}
			}

			if min == max {
				return max
			}
		}
	}
	return 0
}

func mostExpensive(shopList []Shopping) (item []Shopping, err error) {
	// tijelo funkcije
	// pripaziti: kad se može dogoditi pogreška?
	if shopList == nil {
		err = errors.New("No data")
	}

	for _, element := range shopList {
		if len(item) == 0 {
			item = append(item, element)
			continue
		}
		if element.Price == item[0].Price {
			item = append(item, element)
		}
		if element.Price > item[0].Price {
			item = item[:0]
			item = append(item, element)
		}
	}
	return item, err
}
func totalCost(shopList []Shopping) (total int) {
	for _, element := range shopList {
		total += element.Price * element.Quantity
	}
	return total
}
