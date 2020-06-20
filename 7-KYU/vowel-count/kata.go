package main

import "strings"

func VowelCount(str string) (count int) {
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, e := range vowels {
		count += strings.Count(str, e)
	}
	return count
}
