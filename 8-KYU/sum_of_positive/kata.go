package main

func PositiveSum(numbers []int) int {
	num := 0
	for _, number := range numbers {
		if number <= 0 {
			number = 0
		}
		num += number
	}
	return num
}
