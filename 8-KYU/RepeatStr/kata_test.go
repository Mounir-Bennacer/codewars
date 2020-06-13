package main

import "testing"

func TestRepeatStr(t *testing.T) {
	tests := []struct {
		name   string
		inputA int
		inputB string
		output string
	}{
		{name: "test a", inputA: 4, inputB: "a", output: "aaaa"},
		{name: "test hello", inputA: 3, inputB: "hello ", output: "hello hello hello "},
		{name: "test abc", inputA: 2, inputB: "abc", output: "abcabc"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := RepeatStr(test.inputA, test.inputB)
			if got != test.output {
				t.Errorf("%s : got %v expect %s", test.name, got, test.output)
			}
		})
	}
}
