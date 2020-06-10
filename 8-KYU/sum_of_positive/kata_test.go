package main

import "testing"

var data = []struct {
	name   string
	in     []int
	output int
}{
	{"All positive numbers", []int{1, 2, 3, 4, 5}, 15},
	{"One negative number", []int{1, -2, 3, 4, 5}, 13},
	{"Zero number", []int{0}, 0},
	{"All negative numbers", []int{-1, -2, -3, -4, -5}, 0},
	{"Multiple negative numbers", []int{-1, 2, 3, 4, -5}, 9},
}

func TestPositiveSum(t *testing.T) {
	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			got := PositiveSum(tt.in)
			expect := tt.output

			if got != expect {
				t.Fatalf("%v expect %v, got %v", tt.name, expect, got)
			}
		})
	}
}
