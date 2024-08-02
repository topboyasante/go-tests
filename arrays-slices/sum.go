package main

func Sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var result []int
	for _, value := range numbers {
		result = append(result, Sum(value))
	}
	return result
}

func SumAllTails(numbers ...[]int) []int {
	var result []int
	for _, value := range numbers {
		if len(value) == 0 {
			result = append(result, 0)
		} else {
			tails := value[1:]
			result = append(result, Sum(tails))
		}
	}
	return result
}
