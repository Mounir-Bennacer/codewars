package main

import "testing"

func TestMakeNegative(t *testing.T) {
	tests := []struct {
		name          string
		input, output int
	}{
		{name: "negative", input: 41, output: -41},
		{name: "postive", input: -5, output: 5},
		{name: "postive", input: 9, output: -9},
		{name: "postive", input: 649, output: -649},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := MakeNegative(test.input)
			if got != test.output {
				t.Errorf("%s => got %v, expect %v", test.name, got, test.output)
			}
		})
	}
}
