package main

func Sum(numbers []int) int {
	res := 0
	for _, num := range numbers {
		res += num
	}
	return res
}

func SumAll(numbers ...[]int) []int {

	length := len(numbers)
	sums := make([]int, length)

	for i, nums := range numbers {
		sums[i] = Sum(nums)
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, num := range numbers {
		if len(num) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := num[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
