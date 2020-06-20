package main

import "testing"

func TestFindOdd(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{input: []int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}, output: 5},
		{input: []int{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5}, output: -1},
		{input: []int{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5}, output: 5},
		{input: []int{10}, output: 10},
		{input: []int{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1}, output: 10},
		{input: []int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}, output: 1},
	}
	for _, test := range tests {
		t.Run("Should return the approriate number", func(t *testing.T) {
			got := FindOdd(test.input)
			if got != test.output {
				t.Errorf("expected %v , got %v", test.output, got)
			}
		})
	}
}
