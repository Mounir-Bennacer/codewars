package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		ending string
		output bool
	}{
		{name: "empty str", str: "", ending: "", output: true},
		{name: "space str", str: " ", ending: "", output: true},
		{name: "abc", str: "abc", ending: "c", output: true},
		{name: "banana", str: "banana", ending: "ana", output: true},
		{name: "a to z", str: "a", ending: "z", output: false},
		{name: "empty to t", str: "", ending: "t", output: false},
		{name: "$a = $b", str: "$a = $b + 1", ending: "+1", output: false},
		{name: "multiple spaces", str: "    ", ending: "   ", output: true},
		{name: "1oo and 100", str: "1oo", ending: "100", output: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Solution(test.str, test.ending)
			if got != test.output {
				t.Errorf("%v => expect %t, got %t", test.name, test.output, got)
			}
		})
	}
}
