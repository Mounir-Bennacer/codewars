package main

import "testing"

func TestMultiple3And5(t *testing.T) {
	tt := []struct {
		name string
		in   int
		out  int
	}{
		{"Basic cases ", 23, 78},
		{"Zeroes ", 0, 0},
		{"Large numbers ", 23, 9168},
	}
	for _, tc := range tt {
		got := Multiple3And5(tc.in)
		expect := tc.out

		if got != expect {
			t.Errorf("%v => expect %v, got %v", tc.name, expect, got)
		}
	}

}
