package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) (sums []int) {
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}
	return
}

func SumAllTails(slices ...[]int) (sums []int) {
	for _, slice := range slices {
		tail := slice[1:]
		sums = append(sums, Sum(tail))
	}
	return
}
