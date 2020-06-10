package main

import "testing"

var NumToTest = []struct {
	in  int
	out string
}{
	{1, "Odd"},
	{2, "Even"},
	{-1, "Odd"},
	{1, "Odd"},
	{-2, "Even"},
}

func TestEvenOrOdd(t *testing.T) {
	for _, tt := range NumToTest {
		got := EvenOrOdd(tt.in)
		expect := tt.out

		if got != expect {
			t.Errorf("Got %s but expected %s", got, expect)
		}

	}
}
