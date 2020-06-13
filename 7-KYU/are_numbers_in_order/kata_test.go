package main

import "testing"

func TestInAscOrder(t *testing.T) {
	tests := []struct {
		input  []int
		output bool
	}{
		{[]int{1, 2, 4, 7, 19}, true},
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{1, 6, 10, 18, 2, 4, 20}, false},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, false},
	}

	for _, test := range tests {
		t.Run("Should test that result return the correct value", func(t *testing.T) {
			got := InAscOrder(test.input)
			expect := test.output

			if got != expect {
				t.Errorf("got %v, expect %t", got, expect)
			}
		})
	}
}
