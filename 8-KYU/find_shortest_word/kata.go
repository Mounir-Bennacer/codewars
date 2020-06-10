package main

import (
	"strings"
)

// import "strings"

func FindShort(s string) int {
	shortest := len(s)
	for _, word := range strings.Split(s, " ") {
		if len(word) < shortest {
			shortest = len(word)
		}
	}
	return shortest
}
