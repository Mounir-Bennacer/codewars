package main

import "testing"

func TestRemoveChar(t *testing.T) {
	tests := []struct {
		name, input, output string
	}{
		{name: "eloquent", input: "eloquent", output: "loquen"},
		{name: "country", input: "country", output: "ountr"},
		{name: "person", input: "person", output: "erso"},
		{name: "place", input: "place", output: "lac"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := RemoveChar(test.input)
			if got != test.output {
				t.Errorf("%s, got %s, expect %s", test.name, got, test.output)
			}
		})
	}
}
