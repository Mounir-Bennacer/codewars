package main

import "testing"

func TestMonkeyCount(t *testing.T) {
	tests := []struct {
		in  int
		out []int
	}{
		{1, []int{1}},
		{2, []int{1, 2}},
		{3, []int{1, 2, 3}},
		{4, []int{1, 2, 3, 4}},
		{5, []int{1, 2, 3, 4, 5}},
		{6, []int{1, 2, 3, 4, 5, 6}},
		{7, []int{1, 2, 3, 4, 5, 6, 7}},
		{8, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{9, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range tests {
		t.Run("Should count all monkey starting from 1", func(t *testing.T) {
			got := MonkeyCount(test.in)
			if got != test.out {
				t.Errorf("got %v, expect %v", got, test.out)
			}
		})
	}
}
