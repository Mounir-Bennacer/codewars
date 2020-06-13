package main

func MonkeyCount(number int) []int {
	myarray := []int{}
	for i := 1; i <= number; i++ {
		myarray := append(myarray, i)
	}
	return myarray
}
