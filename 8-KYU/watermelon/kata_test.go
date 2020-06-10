package main

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output bool
	}{
		{"4", 4, true},
		{"2", 2, false},
		{"5", 5, false},
		{"88", 88, true},
		{"100", 100, true},
		{"67", 67, false},
		{"90", 90, true},
		{"10", 10, true},
		{"99", 99, false},
		{"32", 32, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Divide(test.input)
			expect := test.output

			if got != expect {
				t.Errorf("%v => got %v, expect %t", test.name, got, expect)
			}
		})
	}
}
