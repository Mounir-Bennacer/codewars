package main

func MonkeyCount(number int) []int {
	myarray := []int{}
	for i := 1; i <= number; i++ {
		_ = append(myarray, i)
	}
	return myarray
}
