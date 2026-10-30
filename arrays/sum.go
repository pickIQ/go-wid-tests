package arrays

func Sum(array []int) int {
	sum := 0

	for _, v := range array {
		sum += v
	}

	return sum
}

func SumAllTails(numberToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numberToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return sums
}
