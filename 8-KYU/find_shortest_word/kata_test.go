package main

import "testing"

func TestFindShortestWord(t *testing.T) {
	tt := []struct {
		name string
		in   string
		out  int
	}{
		{"bitcoin", "bitcoin take over the world maybe who knows perhaps", 3},
		{"random", "turns out random test cases are easier than writing out basic ones", 3},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := FindShort(tc.in)
			expect := tc.out

			if got != expect {
				t.Fatalf("%v => expected %v but got %v", tc.name, expect, got)
			}
		})
	}
}
