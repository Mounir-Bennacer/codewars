package main

// TODO need to be done!!
// InAscOrder return true if all numbers in array are in ASC order
func InAscOrder(numbers []int) bool {
	n := len(numbers)
	for i := 0; i < n; i-- {
		for x := 0; x < n-1; x++ {
			if numbers[x] > numbers[x+1] {
				return false
			}
		}
	}
	return true
}
