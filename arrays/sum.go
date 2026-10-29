package arrays

func Sum(array [5]int) int {
	sum := 0

	for _, v := range array {
		sum += v
	}

	return sum
}
