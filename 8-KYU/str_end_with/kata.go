package main

import (
	"strings"
)

func Solution(str string, ending string) bool {
	return strings.Contains(str, ending)
}
