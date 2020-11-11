package hw1

import "math"

func Counter(numbers []int) map[int]int {
	var pairs = make(map[int]int)
	for _, element := range numbers {
		pairs[element]++
	}
	return pairs
}

func SretanBroj(matrica [][]int) int {
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
