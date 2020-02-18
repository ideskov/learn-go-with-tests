package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	sums := make([]int, len(slices))

	for index, slice := range slices {
		sums[index] = Sum(slice)
	}
	return sums
}
