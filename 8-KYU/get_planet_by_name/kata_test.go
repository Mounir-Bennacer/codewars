package main

import "testing"

var p = []struct {
	in  int
	out string
}{
	{3, "Earth"},
	{2, "Venus"},
	{5, "Jupiter"},
}

func TestGetPlanetName(t *testing.T) {
	for _, tt := range p {
		got := GetPlanetName(tt.in)
		expect := tt.out
		if got != tt.out {
			t.Errorf("Expected %s but got %s", expect, got)
		}
	}
}
