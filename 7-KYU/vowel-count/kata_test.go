package main

import "testing"

func TestVowelCount(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{input: "abracadabra", output: 5},
		{input: "aberlcaidabra", output: 6},
		{input: "abrx", output: 1},
		{input: "abriicaolabria", output: 8},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := VowelCount(test.input)
			if got != test.output {
				t.Errorf("expect %v , got %v", test.output, got)
			}
		})
	}
}
