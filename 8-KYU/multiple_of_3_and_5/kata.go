package main

func Multiple3And5(number int) (result int) {
	// if number == 0 {
	// 	return
	// }
	for i := 0; i < number; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			result = result + i
		}
	}
	return
}
