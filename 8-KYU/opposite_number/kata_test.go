package main

import "testing"

func TestOpposite(t *testing.T) {
	tt := []struct {
		name string
		in   int
		out  int
	}{
		{"negative", -1, 1},
		{"positive", 1, -1},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := Opposite(tc.in)
			expect := tc.out

			if got != expect {
				t.Fatalf("%v => Expect %d, got %d", tc.name, expect, got)
			}
		})
	}
}
