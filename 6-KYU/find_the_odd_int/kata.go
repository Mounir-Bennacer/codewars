package main

func FindOdd(seq []int) (num int) {
	for _, e := range seq {
		num ^= e
	}
	return
}
